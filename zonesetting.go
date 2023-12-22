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
	PrivacyPasses                  *ZoneSettingPrivacyPassService
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
	r.PrivacyPasses = NewZoneSettingPrivacyPassService(opts...)
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
	return
}

// Available settings for your user in relation to a zone.
func (r *ZoneSettingService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSettingsCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit settings for a zone.
func (r *ZoneSettingService) Edit(ctx context.Context, zoneIdentifier string, body ZoneSettingEditParams, opts ...option.RequestOption) (res *ZoneSettingsCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// 0-RTT session resumption enabled for this zone.
type ZeroRtt struct {
	// ID of the zone setting.
	ID ZeroRttID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZeroRttEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the 0-RTT setting.
	Value ZeroRttValue `json:"value"`
	JSON  zeroRttJSON  `json:"-"`
}

// zeroRttJSON contains the JSON metadata for the struct [ZeroRtt]
type zeroRttJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroRtt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroRtt) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZeroRttID string

const (
	ZeroRttID0rtt ZeroRttID = "0rtt"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZeroRttEditable bool

const (
	ZeroRttEditableTrue  ZeroRttEditable = true
	ZeroRttEditableFalse ZeroRttEditable = false
)

// 0-RTT session resumption enabled for this zone.
type ZeroRttParam struct {
	// ID of the zone setting.
	ID param.Field[ZeroRttID] `json:"id"`
	// Value of the 0-RTT setting.
	Value param.Field[ZeroRttValue] `json:"value"`
}

func (r ZeroRttParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZeroRttParam) implementsZoneSettingEditParamsItem() {}

// Value of the 0-RTT setting.
type ZeroRttValue string

const (
	ZeroRttValueOn  ZeroRttValue = "on"
	ZeroRttValueOff ZeroRttValue = "off"
)

type ZoneSettingsCollection struct {
	Errors   []ZoneSettingsCollectionError   `json:"errors"`
	Messages []ZoneSettingsCollectionMessage `json:"messages"`
	Result   []ZoneSettingsCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success bool                       `json:"success"`
	JSON    zoneSettingsCollectionJSON `json:"-"`
}

// zoneSettingsCollectionJSON contains the JSON metadata for the struct
// [ZoneSettingsCollection]
type zoneSettingsCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingsCollectionError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    zoneSettingsCollectionErrorJSON `json:"-"`
}

// zoneSettingsCollectionErrorJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionError]
type zoneSettingsCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSettingsCollectionMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneSettingsCollectionMessageJSON `json:"-"`
}

// zoneSettingsCollectionMessageJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionMessage]
type zoneSettingsCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// 0-RTT session resumption enabled for this zone.
//
// Union satisfied by [ZeroRtt], [ZoneSettingsCollectionResultAdvancedDdos],
// [ZoneSettingsCollectionResultAlwaysOnline],
// [ZoneSettingsCollectionResultAlwaysUseHTTPs],
// [ZoneSettingsCollectionResultAutomaticHTTPsRewrites],
// [ZoneSettingsCollectionResultBrotli],
// [ZoneSettingsCollectionResultBrowserCacheTtl],
// [ZoneSettingsCollectionResultBrowserCheck],
// [ZoneSettingsCollectionResultCacheLevel],
// [ZoneSettingsCollectionResultChallengeTtl],
// [ZoneSettingsCollectionResultCiphers],
// [ZoneSettingsCollectionResultCnameFlattening],
// [ZoneSettingsCollectionResultDevelopmentMode],
// [ZoneSettingsCollectionResultEarlyHints],
// [ZoneSettingsCollectionResultEdgeCacheTtl],
// [ZoneSettingsCollectionResultEmailObfuscation],
// [ZoneSettingsCollectionResultH2Prioritization],
// [ZoneSettingsCollectionResultHotlinkProtection],
// [ZoneSettingsCollectionResultHttp2], [ZoneSettingsCollectionResultHttp3],
// [ZoneSettingsCollectionResultImageResizing],
// [ZoneSettingsCollectionResultIPGeolocation], [ZoneSettingsCollectionResultIpv6],
// [ZoneSettingsCollectionResultMaxUpload],
// [ZoneSettingsCollectionResultMinTlsVersion],
// [ZoneSettingsCollectionResultMinify], [ZoneSettingsCollectionResultMirage],
// [ZoneSettingsCollectionResultMobileRedirect], [ZoneSettingsCollectionResultNel],
// [ZoneSettingsCollectionResultOpportunisticEncryption],
// [ZoneSettingsCollectionResultOpportunisticOnion],
// [ZoneSettingsCollectionResultOrangeToOrange],
// [ZoneSettingsCollectionResultOriginErrorPagePassThru],
// [ZoneSettingsCollectionResultOriginMaxHTTPVersion],
// [ZoneSettingsCollectionResultPolish],
// [ZoneSettingsCollectionResultPrefetchPreload],
// [ZoneSettingsCollectionResultPrivacyPass],
// [ZoneSettingsCollectionResultProxyReadTimeout],
// [ZoneSettingsCollectionResultPseudoIpv4],
// [ZoneSettingsCollectionResultResponseBuffering],
// [ZoneSettingsCollectionResultRocketLoader],
// [ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimization],
// [ZoneSettingsCollectionResultSecurityHeader],
// [ZoneSettingsCollectionResultSecurityLevel],
// [ZoneSettingsCollectionResultServerSideExclude],
// [ZoneSettingsCollectionResultSha1Support],
// [ZoneSettingsCollectionResultSortQueryStringForCache],
// [ZoneSettingsCollectionResultSsl], [ZoneSettingsCollectionResultSslRecommender],
// [ZoneSettingsCollectionResultTls1_2Only], [ZoneSettingsCollectionResultTls1_3],
// [ZoneSettingsCollectionResultTlsClientAuth],
// [ZoneSettingsCollectionResultTrueClientIPHeader],
// [ZoneSettingsCollectionResultWaf], [ZoneSettingsCollectionResultWebp] or
// [ZoneSettingsCollectionResultWebsockets].
type ZoneSettingsCollectionResult interface {
	implementsZoneSettingsCollectionResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneSettingsCollectionResult)(nil)).Elem(), "")
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingsCollectionResultAdvancedDdos struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultAdvancedDdosID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultAdvancedDdosEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingsCollectionResultAdvancedDdosValue `json:"value"`
	JSON  zoneSettingsCollectionResultAdvancedDdosJSON  `json:"-"`
}

// zoneSettingsCollectionResultAdvancedDdosJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultAdvancedDdos]
type zoneSettingsCollectionResultAdvancedDdosJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultAdvancedDdos) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultAdvancedDdos) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultAdvancedDdosID string

const (
	ZoneSettingsCollectionResultAdvancedDdosIDAdvancedDdos ZoneSettingsCollectionResultAdvancedDdosID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultAdvancedDdosEditable bool

const (
	ZoneSettingsCollectionResultAdvancedDdosEditableTrue  ZoneSettingsCollectionResultAdvancedDdosEditable = true
	ZoneSettingsCollectionResultAdvancedDdosEditableFalse ZoneSettingsCollectionResultAdvancedDdosEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingsCollectionResultAdvancedDdosValue string

const (
	ZoneSettingsCollectionResultAdvancedDdosValueOn  ZoneSettingsCollectionResultAdvancedDdosValue = "on"
	ZoneSettingsCollectionResultAdvancedDdosValueOff ZoneSettingsCollectionResultAdvancedDdosValue = "off"
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingsCollectionResultAlwaysOnline struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultAlwaysOnlineID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultAlwaysOnlineEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultAlwaysOnlineValue `json:"value"`
	JSON  zoneSettingsCollectionResultAlwaysOnlineJSON  `json:"-"`
}

// zoneSettingsCollectionResultAlwaysOnlineJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultAlwaysOnline]
type zoneSettingsCollectionResultAlwaysOnlineJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultAlwaysOnline) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultAlwaysOnline) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultAlwaysOnlineID string

const (
	ZoneSettingsCollectionResultAlwaysOnlineIDAlwaysOnline ZoneSettingsCollectionResultAlwaysOnlineID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultAlwaysOnlineEditable bool

const (
	ZoneSettingsCollectionResultAlwaysOnlineEditableTrue  ZoneSettingsCollectionResultAlwaysOnlineEditable = true
	ZoneSettingsCollectionResultAlwaysOnlineEditableFalse ZoneSettingsCollectionResultAlwaysOnlineEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultAlwaysOnlineValue string

const (
	ZoneSettingsCollectionResultAlwaysOnlineValueOn  ZoneSettingsCollectionResultAlwaysOnlineValue = "on"
	ZoneSettingsCollectionResultAlwaysOnlineValueOff ZoneSettingsCollectionResultAlwaysOnlineValue = "off"
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingsCollectionResultAlwaysUseHTTPs struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultAlwaysUseHTTPsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultAlwaysUseHTTPsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultAlwaysUseHTTPsValue `json:"value"`
	JSON  zoneSettingsCollectionResultAlwaysUseHTTPsJSON  `json:"-"`
}

// zoneSettingsCollectionResultAlwaysUseHTTPsJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultAlwaysUseHTTPs]
type zoneSettingsCollectionResultAlwaysUseHTTPsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultAlwaysUseHTTPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultAlwaysUseHTTPs) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultAlwaysUseHTTPsID string

const (
	ZoneSettingsCollectionResultAlwaysUseHTTPsIDAlwaysUseHTTPs ZoneSettingsCollectionResultAlwaysUseHTTPsID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultAlwaysUseHTTPsEditable bool

const (
	ZoneSettingsCollectionResultAlwaysUseHTTPsEditableTrue  ZoneSettingsCollectionResultAlwaysUseHTTPsEditable = true
	ZoneSettingsCollectionResultAlwaysUseHTTPsEditableFalse ZoneSettingsCollectionResultAlwaysUseHTTPsEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultAlwaysUseHTTPsValue string

const (
	ZoneSettingsCollectionResultAlwaysUseHTTPsValueOn  ZoneSettingsCollectionResultAlwaysUseHTTPsValue = "on"
	ZoneSettingsCollectionResultAlwaysUseHTTPsValueOff ZoneSettingsCollectionResultAlwaysUseHTTPsValue = "off"
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingsCollectionResultAutomaticHTTPsRewrites struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultAutomaticHTTPsRewritesID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultAutomaticHTTPsRewritesEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingsCollectionResultAutomaticHTTPsRewritesValue `json:"value"`
	JSON  zoneSettingsCollectionResultAutomaticHTTPsRewritesJSON  `json:"-"`
}

// zoneSettingsCollectionResultAutomaticHTTPsRewritesJSON contains the JSON
// metadata for the struct [ZoneSettingsCollectionResultAutomaticHTTPsRewrites]
type zoneSettingsCollectionResultAutomaticHTTPsRewritesJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultAutomaticHTTPsRewrites) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultAutomaticHTTPsRewrites) implementsZoneSettingsCollectionResult() {
}

// ID of the zone setting.
type ZoneSettingsCollectionResultAutomaticHTTPsRewritesID string

const (
	ZoneSettingsCollectionResultAutomaticHTTPsRewritesIDAutomaticHTTPsRewrites ZoneSettingsCollectionResultAutomaticHTTPsRewritesID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultAutomaticHTTPsRewritesEditable bool

const (
	ZoneSettingsCollectionResultAutomaticHTTPsRewritesEditableTrue  ZoneSettingsCollectionResultAutomaticHTTPsRewritesEditable = true
	ZoneSettingsCollectionResultAutomaticHTTPsRewritesEditableFalse ZoneSettingsCollectionResultAutomaticHTTPsRewritesEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingsCollectionResultAutomaticHTTPsRewritesValue string

const (
	ZoneSettingsCollectionResultAutomaticHTTPsRewritesValueOn  ZoneSettingsCollectionResultAutomaticHTTPsRewritesValue = "on"
	ZoneSettingsCollectionResultAutomaticHTTPsRewritesValueOff ZoneSettingsCollectionResultAutomaticHTTPsRewritesValue = "off"
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingsCollectionResultBrotli struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultBrotliID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultBrotliEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultBrotliValue `json:"value"`
	JSON  zoneSettingsCollectionResultBrotliJSON  `json:"-"`
}

// zoneSettingsCollectionResultBrotliJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultBrotli]
type zoneSettingsCollectionResultBrotliJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultBrotli) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultBrotli) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultBrotliID string

const (
	ZoneSettingsCollectionResultBrotliIDBrotli ZoneSettingsCollectionResultBrotliID = "brotli"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultBrotliEditable bool

const (
	ZoneSettingsCollectionResultBrotliEditableTrue  ZoneSettingsCollectionResultBrotliEditable = true
	ZoneSettingsCollectionResultBrotliEditableFalse ZoneSettingsCollectionResultBrotliEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultBrotliValue string

const (
	ZoneSettingsCollectionResultBrotliValueOff ZoneSettingsCollectionResultBrotliValue = "off"
	ZoneSettingsCollectionResultBrotliValueOn  ZoneSettingsCollectionResultBrotliValue = "on"
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingsCollectionResultBrowserCacheTtl struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultBrowserCacheTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultBrowserCacheTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingsCollectionResultBrowserCacheTtlValue `json:"value"`
	JSON  zoneSettingsCollectionResultBrowserCacheTtlJSON  `json:"-"`
}

// zoneSettingsCollectionResultBrowserCacheTtlJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultBrowserCacheTtl]
type zoneSettingsCollectionResultBrowserCacheTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultBrowserCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultBrowserCacheTtl) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultBrowserCacheTtlID string

const (
	ZoneSettingsCollectionResultBrowserCacheTtlIDBrowserCacheTtl ZoneSettingsCollectionResultBrowserCacheTtlID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultBrowserCacheTtlEditable bool

const (
	ZoneSettingsCollectionResultBrowserCacheTtlEditableTrue  ZoneSettingsCollectionResultBrowserCacheTtlEditable = true
	ZoneSettingsCollectionResultBrowserCacheTtlEditableFalse ZoneSettingsCollectionResultBrowserCacheTtlEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingsCollectionResultBrowserCacheTtlValue float64

const (
	ZoneSettingsCollectionResultBrowserCacheTtlValue0        ZoneSettingsCollectionResultBrowserCacheTtlValue = 0
	ZoneSettingsCollectionResultBrowserCacheTtlValue30       ZoneSettingsCollectionResultBrowserCacheTtlValue = 30
	ZoneSettingsCollectionResultBrowserCacheTtlValue60       ZoneSettingsCollectionResultBrowserCacheTtlValue = 60
	ZoneSettingsCollectionResultBrowserCacheTtlValue120      ZoneSettingsCollectionResultBrowserCacheTtlValue = 120
	ZoneSettingsCollectionResultBrowserCacheTtlValue300      ZoneSettingsCollectionResultBrowserCacheTtlValue = 300
	ZoneSettingsCollectionResultBrowserCacheTtlValue1200     ZoneSettingsCollectionResultBrowserCacheTtlValue = 1200
	ZoneSettingsCollectionResultBrowserCacheTtlValue1800     ZoneSettingsCollectionResultBrowserCacheTtlValue = 1800
	ZoneSettingsCollectionResultBrowserCacheTtlValue3600     ZoneSettingsCollectionResultBrowserCacheTtlValue = 3600
	ZoneSettingsCollectionResultBrowserCacheTtlValue7200     ZoneSettingsCollectionResultBrowserCacheTtlValue = 7200
	ZoneSettingsCollectionResultBrowserCacheTtlValue10800    ZoneSettingsCollectionResultBrowserCacheTtlValue = 10800
	ZoneSettingsCollectionResultBrowserCacheTtlValue14400    ZoneSettingsCollectionResultBrowserCacheTtlValue = 14400
	ZoneSettingsCollectionResultBrowserCacheTtlValue18000    ZoneSettingsCollectionResultBrowserCacheTtlValue = 18000
	ZoneSettingsCollectionResultBrowserCacheTtlValue28800    ZoneSettingsCollectionResultBrowserCacheTtlValue = 28800
	ZoneSettingsCollectionResultBrowserCacheTtlValue43200    ZoneSettingsCollectionResultBrowserCacheTtlValue = 43200
	ZoneSettingsCollectionResultBrowserCacheTtlValue57600    ZoneSettingsCollectionResultBrowserCacheTtlValue = 57600
	ZoneSettingsCollectionResultBrowserCacheTtlValue72000    ZoneSettingsCollectionResultBrowserCacheTtlValue = 72000
	ZoneSettingsCollectionResultBrowserCacheTtlValue86400    ZoneSettingsCollectionResultBrowserCacheTtlValue = 86400
	ZoneSettingsCollectionResultBrowserCacheTtlValue172800   ZoneSettingsCollectionResultBrowserCacheTtlValue = 172800
	ZoneSettingsCollectionResultBrowserCacheTtlValue259200   ZoneSettingsCollectionResultBrowserCacheTtlValue = 259200
	ZoneSettingsCollectionResultBrowserCacheTtlValue345600   ZoneSettingsCollectionResultBrowserCacheTtlValue = 345600
	ZoneSettingsCollectionResultBrowserCacheTtlValue432000   ZoneSettingsCollectionResultBrowserCacheTtlValue = 432000
	ZoneSettingsCollectionResultBrowserCacheTtlValue691200   ZoneSettingsCollectionResultBrowserCacheTtlValue = 691200
	ZoneSettingsCollectionResultBrowserCacheTtlValue1382400  ZoneSettingsCollectionResultBrowserCacheTtlValue = 1382400
	ZoneSettingsCollectionResultBrowserCacheTtlValue2073600  ZoneSettingsCollectionResultBrowserCacheTtlValue = 2073600
	ZoneSettingsCollectionResultBrowserCacheTtlValue2678400  ZoneSettingsCollectionResultBrowserCacheTtlValue = 2678400
	ZoneSettingsCollectionResultBrowserCacheTtlValue5356800  ZoneSettingsCollectionResultBrowserCacheTtlValue = 5356800
	ZoneSettingsCollectionResultBrowserCacheTtlValue16070400 ZoneSettingsCollectionResultBrowserCacheTtlValue = 16070400
	ZoneSettingsCollectionResultBrowserCacheTtlValue31536000 ZoneSettingsCollectionResultBrowserCacheTtlValue = 31536000
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingsCollectionResultBrowserCheck struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultBrowserCheckID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultBrowserCheckEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultBrowserCheckValue `json:"value"`
	JSON  zoneSettingsCollectionResultBrowserCheckJSON  `json:"-"`
}

// zoneSettingsCollectionResultBrowserCheckJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultBrowserCheck]
type zoneSettingsCollectionResultBrowserCheckJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultBrowserCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultBrowserCheck) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultBrowserCheckID string

const (
	ZoneSettingsCollectionResultBrowserCheckIDBrowserCheck ZoneSettingsCollectionResultBrowserCheckID = "browser_check"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultBrowserCheckEditable bool

const (
	ZoneSettingsCollectionResultBrowserCheckEditableTrue  ZoneSettingsCollectionResultBrowserCheckEditable = true
	ZoneSettingsCollectionResultBrowserCheckEditableFalse ZoneSettingsCollectionResultBrowserCheckEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultBrowserCheckValue string

const (
	ZoneSettingsCollectionResultBrowserCheckValueOn  ZoneSettingsCollectionResultBrowserCheckValue = "on"
	ZoneSettingsCollectionResultBrowserCheckValueOff ZoneSettingsCollectionResultBrowserCheckValue = "off"
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingsCollectionResultCacheLevel struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultCacheLevelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultCacheLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultCacheLevelValue `json:"value"`
	JSON  zoneSettingsCollectionResultCacheLevelJSON  `json:"-"`
}

// zoneSettingsCollectionResultCacheLevelJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultCacheLevel]
type zoneSettingsCollectionResultCacheLevelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultCacheLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultCacheLevel) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultCacheLevelID string

const (
	ZoneSettingsCollectionResultCacheLevelIDCacheLevel ZoneSettingsCollectionResultCacheLevelID = "cache_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultCacheLevelEditable bool

const (
	ZoneSettingsCollectionResultCacheLevelEditableTrue  ZoneSettingsCollectionResultCacheLevelEditable = true
	ZoneSettingsCollectionResultCacheLevelEditableFalse ZoneSettingsCollectionResultCacheLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultCacheLevelValue string

const (
	ZoneSettingsCollectionResultCacheLevelValueAggressive ZoneSettingsCollectionResultCacheLevelValue = "aggressive"
	ZoneSettingsCollectionResultCacheLevelValueBasic      ZoneSettingsCollectionResultCacheLevelValue = "basic"
	ZoneSettingsCollectionResultCacheLevelValueSimplified ZoneSettingsCollectionResultCacheLevelValue = "simplified"
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingsCollectionResultChallengeTtl struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultChallengeTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultChallengeTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultChallengeTtlValue `json:"value"`
	JSON  zoneSettingsCollectionResultChallengeTtlJSON  `json:"-"`
}

// zoneSettingsCollectionResultChallengeTtlJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultChallengeTtl]
type zoneSettingsCollectionResultChallengeTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultChallengeTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultChallengeTtl) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultChallengeTtlID string

const (
	ZoneSettingsCollectionResultChallengeTtlIDChallengeTtl ZoneSettingsCollectionResultChallengeTtlID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultChallengeTtlEditable bool

const (
	ZoneSettingsCollectionResultChallengeTtlEditableTrue  ZoneSettingsCollectionResultChallengeTtlEditable = true
	ZoneSettingsCollectionResultChallengeTtlEditableFalse ZoneSettingsCollectionResultChallengeTtlEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultChallengeTtlValue float64

const (
	ZoneSettingsCollectionResultChallengeTtlValue300      ZoneSettingsCollectionResultChallengeTtlValue = 300
	ZoneSettingsCollectionResultChallengeTtlValue900      ZoneSettingsCollectionResultChallengeTtlValue = 900
	ZoneSettingsCollectionResultChallengeTtlValue1800     ZoneSettingsCollectionResultChallengeTtlValue = 1800
	ZoneSettingsCollectionResultChallengeTtlValue2700     ZoneSettingsCollectionResultChallengeTtlValue = 2700
	ZoneSettingsCollectionResultChallengeTtlValue3600     ZoneSettingsCollectionResultChallengeTtlValue = 3600
	ZoneSettingsCollectionResultChallengeTtlValue7200     ZoneSettingsCollectionResultChallengeTtlValue = 7200
	ZoneSettingsCollectionResultChallengeTtlValue10800    ZoneSettingsCollectionResultChallengeTtlValue = 10800
	ZoneSettingsCollectionResultChallengeTtlValue14400    ZoneSettingsCollectionResultChallengeTtlValue = 14400
	ZoneSettingsCollectionResultChallengeTtlValue28800    ZoneSettingsCollectionResultChallengeTtlValue = 28800
	ZoneSettingsCollectionResultChallengeTtlValue57600    ZoneSettingsCollectionResultChallengeTtlValue = 57600
	ZoneSettingsCollectionResultChallengeTtlValue86400    ZoneSettingsCollectionResultChallengeTtlValue = 86400
	ZoneSettingsCollectionResultChallengeTtlValue604800   ZoneSettingsCollectionResultChallengeTtlValue = 604800
	ZoneSettingsCollectionResultChallengeTtlValue2592000  ZoneSettingsCollectionResultChallengeTtlValue = 2592000
	ZoneSettingsCollectionResultChallengeTtlValue31536000 ZoneSettingsCollectionResultChallengeTtlValue = 31536000
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingsCollectionResultCiphers struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultCiphersID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultCiphersEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value []string                                `json:"value"`
	JSON  zoneSettingsCollectionResultCiphersJSON `json:"-"`
}

// zoneSettingsCollectionResultCiphersJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultCiphers]
type zoneSettingsCollectionResultCiphersJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultCiphers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultCiphers) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultCiphersID string

const (
	ZoneSettingsCollectionResultCiphersIDCiphers ZoneSettingsCollectionResultCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultCiphersEditable bool

const (
	ZoneSettingsCollectionResultCiphersEditableTrue  ZoneSettingsCollectionResultCiphersEditable = true
	ZoneSettingsCollectionResultCiphersEditableFalse ZoneSettingsCollectionResultCiphersEditable = false
)

// Whether or not cname flattening is on.
type ZoneSettingsCollectionResultCnameFlattening struct {
	// How to flatten the cname destination.
	ID ZoneSettingsCollectionResultCnameFlatteningID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultCnameFlatteningEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the cname flattening setting.
	Value ZoneSettingsCollectionResultCnameFlatteningValue `json:"value"`
	JSON  zoneSettingsCollectionResultCnameFlatteningJSON  `json:"-"`
}

// zoneSettingsCollectionResultCnameFlatteningJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultCnameFlattening]
type zoneSettingsCollectionResultCnameFlatteningJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultCnameFlattening) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultCnameFlattening) implementsZoneSettingsCollectionResult() {}

// How to flatten the cname destination.
type ZoneSettingsCollectionResultCnameFlatteningID string

const (
	ZoneSettingsCollectionResultCnameFlatteningIDCnameFlattening ZoneSettingsCollectionResultCnameFlatteningID = "cname_flattening"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultCnameFlatteningEditable bool

const (
	ZoneSettingsCollectionResultCnameFlatteningEditableTrue  ZoneSettingsCollectionResultCnameFlatteningEditable = true
	ZoneSettingsCollectionResultCnameFlatteningEditableFalse ZoneSettingsCollectionResultCnameFlatteningEditable = false
)

// Value of the cname flattening setting.
type ZoneSettingsCollectionResultCnameFlatteningValue string

const (
	ZoneSettingsCollectionResultCnameFlatteningValueFlattenAtRoot ZoneSettingsCollectionResultCnameFlatteningValue = "flatten_at_root"
	ZoneSettingsCollectionResultCnameFlatteningValueFlattenAll    ZoneSettingsCollectionResultCnameFlatteningValue = "flatten_all"
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingsCollectionResultDevelopmentMode struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultDevelopmentModeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultDevelopmentModeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The interval (in seconds) from when
	// development mode expires (positive integer) or last expired (negative integer)
	// for the domain. If development mode has never been enabled, this value is false.
	TimeRemaining float64 `json:"time_remaining"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultDevelopmentModeValue `json:"value"`
	JSON  zoneSettingsCollectionResultDevelopmentModeJSON  `json:"-"`
}

// zoneSettingsCollectionResultDevelopmentModeJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultDevelopmentMode]
type zoneSettingsCollectionResultDevelopmentModeJSON struct {
	ID            apijson.Field
	Editable      apijson.Field
	ModifiedOn    apijson.Field
	TimeRemaining apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultDevelopmentMode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultDevelopmentMode) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultDevelopmentModeID string

const (
	ZoneSettingsCollectionResultDevelopmentModeIDDevelopmentMode ZoneSettingsCollectionResultDevelopmentModeID = "development_mode"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultDevelopmentModeEditable bool

const (
	ZoneSettingsCollectionResultDevelopmentModeEditableTrue  ZoneSettingsCollectionResultDevelopmentModeEditable = true
	ZoneSettingsCollectionResultDevelopmentModeEditableFalse ZoneSettingsCollectionResultDevelopmentModeEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultDevelopmentModeValue string

const (
	ZoneSettingsCollectionResultDevelopmentModeValueOn  ZoneSettingsCollectionResultDevelopmentModeValue = "on"
	ZoneSettingsCollectionResultDevelopmentModeValueOff ZoneSettingsCollectionResultDevelopmentModeValue = "off"
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingsCollectionResultEarlyHints struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultEarlyHintsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultEarlyHintsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultEarlyHintsValue `json:"value"`
	JSON  zoneSettingsCollectionResultEarlyHintsJSON  `json:"-"`
}

// zoneSettingsCollectionResultEarlyHintsJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultEarlyHints]
type zoneSettingsCollectionResultEarlyHintsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultEarlyHints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultEarlyHints) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultEarlyHintsID string

const (
	ZoneSettingsCollectionResultEarlyHintsIDEarlyHints ZoneSettingsCollectionResultEarlyHintsID = "early_hints"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultEarlyHintsEditable bool

const (
	ZoneSettingsCollectionResultEarlyHintsEditableTrue  ZoneSettingsCollectionResultEarlyHintsEditable = true
	ZoneSettingsCollectionResultEarlyHintsEditableFalse ZoneSettingsCollectionResultEarlyHintsEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultEarlyHintsValue string

const (
	ZoneSettingsCollectionResultEarlyHintsValueOn  ZoneSettingsCollectionResultEarlyHintsValue = "on"
	ZoneSettingsCollectionResultEarlyHintsValueOff ZoneSettingsCollectionResultEarlyHintsValue = "off"
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingsCollectionResultEdgeCacheTtl struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultEdgeCacheTtlID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultEdgeCacheTtlEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value ZoneSettingsCollectionResultEdgeCacheTtlValue `json:"value"`
	JSON  zoneSettingsCollectionResultEdgeCacheTtlJSON  `json:"-"`
}

// zoneSettingsCollectionResultEdgeCacheTtlJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultEdgeCacheTtl]
type zoneSettingsCollectionResultEdgeCacheTtlJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultEdgeCacheTtl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultEdgeCacheTtl) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultEdgeCacheTtlID string

const (
	ZoneSettingsCollectionResultEdgeCacheTtlIDEdgeCacheTtl ZoneSettingsCollectionResultEdgeCacheTtlID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultEdgeCacheTtlEditable bool

const (
	ZoneSettingsCollectionResultEdgeCacheTtlEditableTrue  ZoneSettingsCollectionResultEdgeCacheTtlEditable = true
	ZoneSettingsCollectionResultEdgeCacheTtlEditableFalse ZoneSettingsCollectionResultEdgeCacheTtlEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingsCollectionResultEdgeCacheTtlValue float64

const (
	ZoneSettingsCollectionResultEdgeCacheTtlValue30     ZoneSettingsCollectionResultEdgeCacheTtlValue = 30
	ZoneSettingsCollectionResultEdgeCacheTtlValue60     ZoneSettingsCollectionResultEdgeCacheTtlValue = 60
	ZoneSettingsCollectionResultEdgeCacheTtlValue300    ZoneSettingsCollectionResultEdgeCacheTtlValue = 300
	ZoneSettingsCollectionResultEdgeCacheTtlValue1200   ZoneSettingsCollectionResultEdgeCacheTtlValue = 1200
	ZoneSettingsCollectionResultEdgeCacheTtlValue1800   ZoneSettingsCollectionResultEdgeCacheTtlValue = 1800
	ZoneSettingsCollectionResultEdgeCacheTtlValue3600   ZoneSettingsCollectionResultEdgeCacheTtlValue = 3600
	ZoneSettingsCollectionResultEdgeCacheTtlValue7200   ZoneSettingsCollectionResultEdgeCacheTtlValue = 7200
	ZoneSettingsCollectionResultEdgeCacheTtlValue10800  ZoneSettingsCollectionResultEdgeCacheTtlValue = 10800
	ZoneSettingsCollectionResultEdgeCacheTtlValue14400  ZoneSettingsCollectionResultEdgeCacheTtlValue = 14400
	ZoneSettingsCollectionResultEdgeCacheTtlValue18000  ZoneSettingsCollectionResultEdgeCacheTtlValue = 18000
	ZoneSettingsCollectionResultEdgeCacheTtlValue28800  ZoneSettingsCollectionResultEdgeCacheTtlValue = 28800
	ZoneSettingsCollectionResultEdgeCacheTtlValue43200  ZoneSettingsCollectionResultEdgeCacheTtlValue = 43200
	ZoneSettingsCollectionResultEdgeCacheTtlValue57600  ZoneSettingsCollectionResultEdgeCacheTtlValue = 57600
	ZoneSettingsCollectionResultEdgeCacheTtlValue72000  ZoneSettingsCollectionResultEdgeCacheTtlValue = 72000
	ZoneSettingsCollectionResultEdgeCacheTtlValue86400  ZoneSettingsCollectionResultEdgeCacheTtlValue = 86400
	ZoneSettingsCollectionResultEdgeCacheTtlValue172800 ZoneSettingsCollectionResultEdgeCacheTtlValue = 172800
	ZoneSettingsCollectionResultEdgeCacheTtlValue259200 ZoneSettingsCollectionResultEdgeCacheTtlValue = 259200
	ZoneSettingsCollectionResultEdgeCacheTtlValue345600 ZoneSettingsCollectionResultEdgeCacheTtlValue = 345600
	ZoneSettingsCollectionResultEdgeCacheTtlValue432000 ZoneSettingsCollectionResultEdgeCacheTtlValue = 432000
	ZoneSettingsCollectionResultEdgeCacheTtlValue518400 ZoneSettingsCollectionResultEdgeCacheTtlValue = 518400
	ZoneSettingsCollectionResultEdgeCacheTtlValue604800 ZoneSettingsCollectionResultEdgeCacheTtlValue = 604800
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingsCollectionResultEmailObfuscation struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultEmailObfuscationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultEmailObfuscationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultEmailObfuscationValue `json:"value"`
	JSON  zoneSettingsCollectionResultEmailObfuscationJSON  `json:"-"`
}

// zoneSettingsCollectionResultEmailObfuscationJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultEmailObfuscation]
type zoneSettingsCollectionResultEmailObfuscationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultEmailObfuscation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultEmailObfuscation) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultEmailObfuscationID string

const (
	ZoneSettingsCollectionResultEmailObfuscationIDEmailObfuscation ZoneSettingsCollectionResultEmailObfuscationID = "email_obfuscation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultEmailObfuscationEditable bool

const (
	ZoneSettingsCollectionResultEmailObfuscationEditableTrue  ZoneSettingsCollectionResultEmailObfuscationEditable = true
	ZoneSettingsCollectionResultEmailObfuscationEditableFalse ZoneSettingsCollectionResultEmailObfuscationEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultEmailObfuscationValue string

const (
	ZoneSettingsCollectionResultEmailObfuscationValueOn  ZoneSettingsCollectionResultEmailObfuscationValue = "on"
	ZoneSettingsCollectionResultEmailObfuscationValueOff ZoneSettingsCollectionResultEmailObfuscationValue = "off"
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingsCollectionResultH2Prioritization struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultH2PrioritizationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultH2PrioritizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultH2PrioritizationValue `json:"value"`
	JSON  zoneSettingsCollectionResultH2PrioritizationJSON  `json:"-"`
}

// zoneSettingsCollectionResultH2PrioritizationJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultH2Prioritization]
type zoneSettingsCollectionResultH2PrioritizationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultH2Prioritization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultH2Prioritization) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultH2PrioritizationID string

const (
	ZoneSettingsCollectionResultH2PrioritizationIDH2Prioritization ZoneSettingsCollectionResultH2PrioritizationID = "h2_prioritization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultH2PrioritizationEditable bool

const (
	ZoneSettingsCollectionResultH2PrioritizationEditableTrue  ZoneSettingsCollectionResultH2PrioritizationEditable = true
	ZoneSettingsCollectionResultH2PrioritizationEditableFalse ZoneSettingsCollectionResultH2PrioritizationEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultH2PrioritizationValue string

const (
	ZoneSettingsCollectionResultH2PrioritizationValueOn     ZoneSettingsCollectionResultH2PrioritizationValue = "on"
	ZoneSettingsCollectionResultH2PrioritizationValueOff    ZoneSettingsCollectionResultH2PrioritizationValue = "off"
	ZoneSettingsCollectionResultH2PrioritizationValueCustom ZoneSettingsCollectionResultH2PrioritizationValue = "custom"
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingsCollectionResultHotlinkProtection struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultHotlinkProtectionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultHotlinkProtectionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultHotlinkProtectionValue `json:"value"`
	JSON  zoneSettingsCollectionResultHotlinkProtectionJSON  `json:"-"`
}

// zoneSettingsCollectionResultHotlinkProtectionJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultHotlinkProtection]
type zoneSettingsCollectionResultHotlinkProtectionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultHotlinkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultHotlinkProtection) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultHotlinkProtectionID string

const (
	ZoneSettingsCollectionResultHotlinkProtectionIDHotlinkProtection ZoneSettingsCollectionResultHotlinkProtectionID = "hotlink_protection"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultHotlinkProtectionEditable bool

const (
	ZoneSettingsCollectionResultHotlinkProtectionEditableTrue  ZoneSettingsCollectionResultHotlinkProtectionEditable = true
	ZoneSettingsCollectionResultHotlinkProtectionEditableFalse ZoneSettingsCollectionResultHotlinkProtectionEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultHotlinkProtectionValue string

const (
	ZoneSettingsCollectionResultHotlinkProtectionValueOn  ZoneSettingsCollectionResultHotlinkProtectionValue = "on"
	ZoneSettingsCollectionResultHotlinkProtectionValueOff ZoneSettingsCollectionResultHotlinkProtectionValue = "off"
)

// HTTP2 enabled for this zone.
type ZoneSettingsCollectionResultHttp2 struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultHttp2ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultHttp2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingsCollectionResultHttp2Value `json:"value"`
	JSON  zoneSettingsCollectionResultHttp2JSON  `json:"-"`
}

// zoneSettingsCollectionResultHttp2JSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultHttp2]
type zoneSettingsCollectionResultHttp2JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultHttp2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultHttp2) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultHttp2ID string

const (
	ZoneSettingsCollectionResultHttp2IDHttp2 ZoneSettingsCollectionResultHttp2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultHttp2Editable bool

const (
	ZoneSettingsCollectionResultHttp2EditableTrue  ZoneSettingsCollectionResultHttp2Editable = true
	ZoneSettingsCollectionResultHttp2EditableFalse ZoneSettingsCollectionResultHttp2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingsCollectionResultHttp2Value string

const (
	ZoneSettingsCollectionResultHttp2ValueOn  ZoneSettingsCollectionResultHttp2Value = "on"
	ZoneSettingsCollectionResultHttp2ValueOff ZoneSettingsCollectionResultHttp2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingsCollectionResultHttp3 struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultHttp3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultHttp3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingsCollectionResultHttp3Value `json:"value"`
	JSON  zoneSettingsCollectionResultHttp3JSON  `json:"-"`
}

// zoneSettingsCollectionResultHttp3JSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultHttp3]
type zoneSettingsCollectionResultHttp3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultHttp3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultHttp3) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultHttp3ID string

const (
	ZoneSettingsCollectionResultHttp3IDHttp3 ZoneSettingsCollectionResultHttp3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultHttp3Editable bool

const (
	ZoneSettingsCollectionResultHttp3EditableTrue  ZoneSettingsCollectionResultHttp3Editable = true
	ZoneSettingsCollectionResultHttp3EditableFalse ZoneSettingsCollectionResultHttp3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingsCollectionResultHttp3Value string

const (
	ZoneSettingsCollectionResultHttp3ValueOn  ZoneSettingsCollectionResultHttp3Value = "on"
	ZoneSettingsCollectionResultHttp3ValueOff ZoneSettingsCollectionResultHttp3Value = "off"
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingsCollectionResultImageResizing struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultImageResizingID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultImageResizingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value ZoneSettingsCollectionResultImageResizingValue `json:"value"`
	JSON  zoneSettingsCollectionResultImageResizingJSON  `json:"-"`
}

// zoneSettingsCollectionResultImageResizingJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultImageResizing]
type zoneSettingsCollectionResultImageResizingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultImageResizing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultImageResizing) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultImageResizingID string

const (
	ZoneSettingsCollectionResultImageResizingIDImageResizing ZoneSettingsCollectionResultImageResizingID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultImageResizingEditable bool

const (
	ZoneSettingsCollectionResultImageResizingEditableTrue  ZoneSettingsCollectionResultImageResizingEditable = true
	ZoneSettingsCollectionResultImageResizingEditableFalse ZoneSettingsCollectionResultImageResizingEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingsCollectionResultImageResizingValue string

const (
	ZoneSettingsCollectionResultImageResizingValueOn   ZoneSettingsCollectionResultImageResizingValue = "on"
	ZoneSettingsCollectionResultImageResizingValueOff  ZoneSettingsCollectionResultImageResizingValue = "off"
	ZoneSettingsCollectionResultImageResizingValueOpen ZoneSettingsCollectionResultImageResizingValue = "open"
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingsCollectionResultIPGeolocation struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultIPGeolocationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultIPGeolocationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultIPGeolocationValue `json:"value"`
	JSON  zoneSettingsCollectionResultIPGeolocationJSON  `json:"-"`
}

// zoneSettingsCollectionResultIPGeolocationJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultIPGeolocation]
type zoneSettingsCollectionResultIPGeolocationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultIPGeolocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultIPGeolocation) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultIPGeolocationID string

const (
	ZoneSettingsCollectionResultIPGeolocationIDIPGeolocation ZoneSettingsCollectionResultIPGeolocationID = "ip_geolocation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultIPGeolocationEditable bool

const (
	ZoneSettingsCollectionResultIPGeolocationEditableTrue  ZoneSettingsCollectionResultIPGeolocationEditable = true
	ZoneSettingsCollectionResultIPGeolocationEditableFalse ZoneSettingsCollectionResultIPGeolocationEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultIPGeolocationValue string

const (
	ZoneSettingsCollectionResultIPGeolocationValueOn  ZoneSettingsCollectionResultIPGeolocationValue = "on"
	ZoneSettingsCollectionResultIPGeolocationValueOff ZoneSettingsCollectionResultIPGeolocationValue = "off"
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingsCollectionResultIpv6 struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultIpv6ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultIpv6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultIpv6Value `json:"value"`
	JSON  zoneSettingsCollectionResultIpv6JSON  `json:"-"`
}

// zoneSettingsCollectionResultIpv6JSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultIpv6]
type zoneSettingsCollectionResultIpv6JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultIpv6) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultIpv6ID string

const (
	ZoneSettingsCollectionResultIpv6IDIpv6 ZoneSettingsCollectionResultIpv6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultIpv6Editable bool

const (
	ZoneSettingsCollectionResultIpv6EditableTrue  ZoneSettingsCollectionResultIpv6Editable = true
	ZoneSettingsCollectionResultIpv6EditableFalse ZoneSettingsCollectionResultIpv6Editable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultIpv6Value string

const (
	ZoneSettingsCollectionResultIpv6ValueOff ZoneSettingsCollectionResultIpv6Value = "off"
	ZoneSettingsCollectionResultIpv6ValueOn  ZoneSettingsCollectionResultIpv6Value = "on"
)

// Maximum size of an allowable upload.
type ZoneSettingsCollectionResultMaxUpload struct {
	// identifier of the zone setting.
	ID ZoneSettingsCollectionResultMaxUploadID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultMaxUploadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value ZoneSettingsCollectionResultMaxUploadValue `json:"value"`
	JSON  zoneSettingsCollectionResultMaxUploadJSON  `json:"-"`
}

// zoneSettingsCollectionResultMaxUploadJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultMaxUpload]
type zoneSettingsCollectionResultMaxUploadJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMaxUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultMaxUpload) implementsZoneSettingsCollectionResult() {}

// identifier of the zone setting.
type ZoneSettingsCollectionResultMaxUploadID string

const (
	ZoneSettingsCollectionResultMaxUploadIDMaxUpload ZoneSettingsCollectionResultMaxUploadID = "max_upload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultMaxUploadEditable bool

const (
	ZoneSettingsCollectionResultMaxUploadEditableTrue  ZoneSettingsCollectionResultMaxUploadEditable = true
	ZoneSettingsCollectionResultMaxUploadEditableFalse ZoneSettingsCollectionResultMaxUploadEditable = false
)

// Value of the zone setting. Notes: The size depends on the plan level of the
// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
type ZoneSettingsCollectionResultMaxUploadValue float64

const (
	ZoneSettingsCollectionResultMaxUploadValue100 ZoneSettingsCollectionResultMaxUploadValue = 100
	ZoneSettingsCollectionResultMaxUploadValue200 ZoneSettingsCollectionResultMaxUploadValue = 200
	ZoneSettingsCollectionResultMaxUploadValue500 ZoneSettingsCollectionResultMaxUploadValue = 500
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingsCollectionResultMinTlsVersion struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultMinTlsVersionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultMinTlsVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultMinTlsVersionValue `json:"value"`
	JSON  zoneSettingsCollectionResultMinTlsVersionJSON  `json:"-"`
}

// zoneSettingsCollectionResultMinTlsVersionJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultMinTlsVersion]
type zoneSettingsCollectionResultMinTlsVersionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMinTlsVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultMinTlsVersion) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultMinTlsVersionID string

const (
	ZoneSettingsCollectionResultMinTlsVersionIDMinTlsVersion ZoneSettingsCollectionResultMinTlsVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultMinTlsVersionEditable bool

const (
	ZoneSettingsCollectionResultMinTlsVersionEditableTrue  ZoneSettingsCollectionResultMinTlsVersionEditable = true
	ZoneSettingsCollectionResultMinTlsVersionEditableFalse ZoneSettingsCollectionResultMinTlsVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultMinTlsVersionValue string

const (
	ZoneSettingsCollectionResultMinTlsVersionValue1_0 ZoneSettingsCollectionResultMinTlsVersionValue = "1.0"
	ZoneSettingsCollectionResultMinTlsVersionValue1_1 ZoneSettingsCollectionResultMinTlsVersionValue = "1.1"
	ZoneSettingsCollectionResultMinTlsVersionValue1_2 ZoneSettingsCollectionResultMinTlsVersionValue = "1.2"
	ZoneSettingsCollectionResultMinTlsVersionValue1_3 ZoneSettingsCollectionResultMinTlsVersionValue = "1.3"
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingsCollectionResultMinify struct {
	// Zone setting identifier.
	ID ZoneSettingsCollectionResultMinifyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultMinifyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultMinifyValue `json:"value"`
	JSON  zoneSettingsCollectionResultMinifyJSON  `json:"-"`
}

// zoneSettingsCollectionResultMinifyJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultMinify]
type zoneSettingsCollectionResultMinifyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMinify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultMinify) implementsZoneSettingsCollectionResult() {}

// Zone setting identifier.
type ZoneSettingsCollectionResultMinifyID string

const (
	ZoneSettingsCollectionResultMinifyIDMinify ZoneSettingsCollectionResultMinifyID = "minify"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultMinifyEditable bool

const (
	ZoneSettingsCollectionResultMinifyEditableTrue  ZoneSettingsCollectionResultMinifyEditable = true
	ZoneSettingsCollectionResultMinifyEditableFalse ZoneSettingsCollectionResultMinifyEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css ZoneSettingsCollectionResultMinifyValueCss `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML ZoneSettingsCollectionResultMinifyValueHTML `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js   ZoneSettingsCollectionResultMinifyValueJs   `json:"js"`
	JSON zoneSettingsCollectionResultMinifyValueJSON `json:"-"`
}

// zoneSettingsCollectionResultMinifyValueJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultMinifyValue]
type zoneSettingsCollectionResultMinifyValueJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMinifyValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingsCollectionResultMinifyValueCss string

const (
	ZoneSettingsCollectionResultMinifyValueCssOn  ZoneSettingsCollectionResultMinifyValueCss = "on"
	ZoneSettingsCollectionResultMinifyValueCssOff ZoneSettingsCollectionResultMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingsCollectionResultMinifyValueHTML string

const (
	ZoneSettingsCollectionResultMinifyValueHTMLOn  ZoneSettingsCollectionResultMinifyValueHTML = "on"
	ZoneSettingsCollectionResultMinifyValueHTMLOff ZoneSettingsCollectionResultMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingsCollectionResultMinifyValueJs string

const (
	ZoneSettingsCollectionResultMinifyValueJsOn  ZoneSettingsCollectionResultMinifyValueJs = "on"
	ZoneSettingsCollectionResultMinifyValueJsOff ZoneSettingsCollectionResultMinifyValueJs = "off"
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingsCollectionResultMirage struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultMirageID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultMirageEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultMirageValue `json:"value"`
	JSON  zoneSettingsCollectionResultMirageJSON  `json:"-"`
}

// zoneSettingsCollectionResultMirageJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultMirage]
type zoneSettingsCollectionResultMirageJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMirage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultMirage) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultMirageID string

const (
	ZoneSettingsCollectionResultMirageIDMirage ZoneSettingsCollectionResultMirageID = "mirage"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultMirageEditable bool

const (
	ZoneSettingsCollectionResultMirageEditableTrue  ZoneSettingsCollectionResultMirageEditable = true
	ZoneSettingsCollectionResultMirageEditableFalse ZoneSettingsCollectionResultMirageEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultMirageValue string

const (
	ZoneSettingsCollectionResultMirageValueOn  ZoneSettingsCollectionResultMirageValue = "on"
	ZoneSettingsCollectionResultMirageValueOff ZoneSettingsCollectionResultMirageValue = "off"
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingsCollectionResultMobileRedirect struct {
	// Identifier of the zone setting.
	ID ZoneSettingsCollectionResultMobileRedirectID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultMobileRedirectEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultMobileRedirectValue `json:"value"`
	JSON  zoneSettingsCollectionResultMobileRedirectJSON  `json:"-"`
}

// zoneSettingsCollectionResultMobileRedirectJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultMobileRedirect]
type zoneSettingsCollectionResultMobileRedirectJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMobileRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultMobileRedirect) implementsZoneSettingsCollectionResult() {}

// Identifier of the zone setting.
type ZoneSettingsCollectionResultMobileRedirectID string

const (
	ZoneSettingsCollectionResultMobileRedirectIDMobileRedirect ZoneSettingsCollectionResultMobileRedirectID = "mobile_redirect"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultMobileRedirectEditable bool

const (
	ZoneSettingsCollectionResultMobileRedirectEditableTrue  ZoneSettingsCollectionResultMobileRedirectEditable = true
	ZoneSettingsCollectionResultMobileRedirectEditableFalse ZoneSettingsCollectionResultMobileRedirectEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain string `json:"mobile_subdomain,nullable"`
	// Whether or not mobile redirect is enabled.
	Status ZoneSettingsCollectionResultMobileRedirectValueStatus `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri bool                                                `json:"strip_uri"`
	JSON     zoneSettingsCollectionResultMobileRedirectValueJSON `json:"-"`
}

// zoneSettingsCollectionResultMobileRedirectValueJSON contains the JSON metadata
// for the struct [ZoneSettingsCollectionResultMobileRedirectValue]
type zoneSettingsCollectionResultMobileRedirectValueJSON struct {
	MobileSubdomain apijson.Field
	Status          apijson.Field
	StripUri        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultMobileRedirectValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingsCollectionResultMobileRedirectValueStatus string

const (
	ZoneSettingsCollectionResultMobileRedirectValueStatusOn  ZoneSettingsCollectionResultMobileRedirectValueStatus = "on"
	ZoneSettingsCollectionResultMobileRedirectValueStatusOff ZoneSettingsCollectionResultMobileRedirectValueStatus = "off"
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingsCollectionResultNel struct {
	// Zone setting identifier.
	ID ZoneSettingsCollectionResultNelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultNelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultNelValue `json:"value"`
	JSON  zoneSettingsCollectionResultNelJSON  `json:"-"`
}

// zoneSettingsCollectionResultNelJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultNel]
type zoneSettingsCollectionResultNelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultNel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultNel) implementsZoneSettingsCollectionResult() {}

// Zone setting identifier.
type ZoneSettingsCollectionResultNelID string

const (
	ZoneSettingsCollectionResultNelIDNel ZoneSettingsCollectionResultNelID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultNelEditable bool

const (
	ZoneSettingsCollectionResultNelEditableTrue  ZoneSettingsCollectionResultNelEditable = true
	ZoneSettingsCollectionResultNelEditableFalse ZoneSettingsCollectionResultNelEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultNelValue struct {
	Enabled bool                                     `json:"enabled"`
	JSON    zoneSettingsCollectionResultNelValueJSON `json:"-"`
}

// zoneSettingsCollectionResultNelValueJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultNelValue]
type zoneSettingsCollectionResultNelValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultNelValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingsCollectionResultOpportunisticEncryption struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultOpportunisticEncryptionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultOpportunisticEncryptionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingsCollectionResultOpportunisticEncryptionValue `json:"value"`
	JSON  zoneSettingsCollectionResultOpportunisticEncryptionJSON  `json:"-"`
}

// zoneSettingsCollectionResultOpportunisticEncryptionJSON contains the JSON
// metadata for the struct [ZoneSettingsCollectionResultOpportunisticEncryption]
type zoneSettingsCollectionResultOpportunisticEncryptionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultOpportunisticEncryption) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultOpportunisticEncryption) implementsZoneSettingsCollectionResult() {
}

// ID of the zone setting.
type ZoneSettingsCollectionResultOpportunisticEncryptionID string

const (
	ZoneSettingsCollectionResultOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingsCollectionResultOpportunisticEncryptionID = "opportunistic_encryption"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultOpportunisticEncryptionEditable bool

const (
	ZoneSettingsCollectionResultOpportunisticEncryptionEditableTrue  ZoneSettingsCollectionResultOpportunisticEncryptionEditable = true
	ZoneSettingsCollectionResultOpportunisticEncryptionEditableFalse ZoneSettingsCollectionResultOpportunisticEncryptionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingsCollectionResultOpportunisticEncryptionValue string

const (
	ZoneSettingsCollectionResultOpportunisticEncryptionValueOn  ZoneSettingsCollectionResultOpportunisticEncryptionValue = "on"
	ZoneSettingsCollectionResultOpportunisticEncryptionValueOff ZoneSettingsCollectionResultOpportunisticEncryptionValue = "off"
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingsCollectionResultOpportunisticOnion struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultOpportunisticOnionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultOpportunisticOnionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingsCollectionResultOpportunisticOnionValue `json:"value"`
	JSON  zoneSettingsCollectionResultOpportunisticOnionJSON  `json:"-"`
}

// zoneSettingsCollectionResultOpportunisticOnionJSON contains the JSON metadata
// for the struct [ZoneSettingsCollectionResultOpportunisticOnion]
type zoneSettingsCollectionResultOpportunisticOnionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultOpportunisticOnion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultOpportunisticOnion) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultOpportunisticOnionID string

const (
	ZoneSettingsCollectionResultOpportunisticOnionIDOpportunisticOnion ZoneSettingsCollectionResultOpportunisticOnionID = "opportunistic_onion"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultOpportunisticOnionEditable bool

const (
	ZoneSettingsCollectionResultOpportunisticOnionEditableTrue  ZoneSettingsCollectionResultOpportunisticOnionEditable = true
	ZoneSettingsCollectionResultOpportunisticOnionEditableFalse ZoneSettingsCollectionResultOpportunisticOnionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingsCollectionResultOpportunisticOnionValue string

const (
	ZoneSettingsCollectionResultOpportunisticOnionValueOn  ZoneSettingsCollectionResultOpportunisticOnionValue = "on"
	ZoneSettingsCollectionResultOpportunisticOnionValueOff ZoneSettingsCollectionResultOpportunisticOnionValue = "off"
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingsCollectionResultOrangeToOrange struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultOrangeToOrangeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultOrangeToOrangeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultOrangeToOrangeValue `json:"value"`
	JSON  zoneSettingsCollectionResultOrangeToOrangeJSON  `json:"-"`
}

// zoneSettingsCollectionResultOrangeToOrangeJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultOrangeToOrange]
type zoneSettingsCollectionResultOrangeToOrangeJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultOrangeToOrange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultOrangeToOrange) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultOrangeToOrangeID string

const (
	ZoneSettingsCollectionResultOrangeToOrangeIDOrangeToOrange ZoneSettingsCollectionResultOrangeToOrangeID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultOrangeToOrangeEditable bool

const (
	ZoneSettingsCollectionResultOrangeToOrangeEditableTrue  ZoneSettingsCollectionResultOrangeToOrangeEditable = true
	ZoneSettingsCollectionResultOrangeToOrangeEditableFalse ZoneSettingsCollectionResultOrangeToOrangeEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultOrangeToOrangeValue string

const (
	ZoneSettingsCollectionResultOrangeToOrangeValueOn  ZoneSettingsCollectionResultOrangeToOrangeValue = "on"
	ZoneSettingsCollectionResultOrangeToOrangeValueOff ZoneSettingsCollectionResultOrangeToOrangeValue = "off"
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingsCollectionResultOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultOriginErrorPagePassThruID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultOriginErrorPagePassThruEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultOriginErrorPagePassThruValue `json:"value"`
	JSON  zoneSettingsCollectionResultOriginErrorPagePassThruJSON  `json:"-"`
}

// zoneSettingsCollectionResultOriginErrorPagePassThruJSON contains the JSON
// metadata for the struct [ZoneSettingsCollectionResultOriginErrorPagePassThru]
type zoneSettingsCollectionResultOriginErrorPagePassThruJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultOriginErrorPagePassThru) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultOriginErrorPagePassThru) implementsZoneSettingsCollectionResult() {
}

// ID of the zone setting.
type ZoneSettingsCollectionResultOriginErrorPagePassThruID string

const (
	ZoneSettingsCollectionResultOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingsCollectionResultOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultOriginErrorPagePassThruEditable bool

const (
	ZoneSettingsCollectionResultOriginErrorPagePassThruEditableTrue  ZoneSettingsCollectionResultOriginErrorPagePassThruEditable = true
	ZoneSettingsCollectionResultOriginErrorPagePassThruEditableFalse ZoneSettingsCollectionResultOriginErrorPagePassThruEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultOriginErrorPagePassThruValue string

const (
	ZoneSettingsCollectionResultOriginErrorPagePassThruValueOn  ZoneSettingsCollectionResultOriginErrorPagePassThruValue = "on"
	ZoneSettingsCollectionResultOriginErrorPagePassThruValueOff ZoneSettingsCollectionResultOriginErrorPagePassThruValue = "off"
)

type ZoneSettingsCollectionResultOriginMaxHTTPVersion struct {
	// Identifier of the zone setting.
	ID ZoneSettingsCollectionResultOriginMaxHTTPVersionID `json:"id,required"`
	// last time this setting was modified.
	ModifiedOn time.Time                                            `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingsCollectionResultOriginMaxHTTPVersionJSON `json:"-"`
}

// zoneSettingsCollectionResultOriginMaxHTTPVersionJSON contains the JSON metadata
// for the struct [ZoneSettingsCollectionResultOriginMaxHTTPVersion]
type zoneSettingsCollectionResultOriginMaxHTTPVersionJSON struct {
	ID          apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultOriginMaxHTTPVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultOriginMaxHTTPVersion) implementsZoneSettingsCollectionResult() {}

// Identifier of the zone setting.
type ZoneSettingsCollectionResultOriginMaxHTTPVersionID string

const (
	ZoneSettingsCollectionResultOriginMaxHTTPVersionIDOriginMaxHTTPVersion ZoneSettingsCollectionResultOriginMaxHTTPVersionID = "origin_max_http_version"
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingsCollectionResultPolish struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultPolishID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultPolishEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultPolishValue `json:"value"`
	JSON  zoneSettingsCollectionResultPolishJSON  `json:"-"`
}

// zoneSettingsCollectionResultPolishJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultPolish]
type zoneSettingsCollectionResultPolishJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultPolish) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultPolish) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultPolishID string

const (
	ZoneSettingsCollectionResultPolishIDPolish ZoneSettingsCollectionResultPolishID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultPolishEditable bool

const (
	ZoneSettingsCollectionResultPolishEditableTrue  ZoneSettingsCollectionResultPolishEditable = true
	ZoneSettingsCollectionResultPolishEditableFalse ZoneSettingsCollectionResultPolishEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultPolishValue string

const (
	ZoneSettingsCollectionResultPolishValueOff      ZoneSettingsCollectionResultPolishValue = "off"
	ZoneSettingsCollectionResultPolishValueLossless ZoneSettingsCollectionResultPolishValue = "lossless"
	ZoneSettingsCollectionResultPolishValueLossy    ZoneSettingsCollectionResultPolishValue = "lossy"
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingsCollectionResultPrefetchPreload struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultPrefetchPreloadID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultPrefetchPreloadEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultPrefetchPreloadValue `json:"value"`
	JSON  zoneSettingsCollectionResultPrefetchPreloadJSON  `json:"-"`
}

// zoneSettingsCollectionResultPrefetchPreloadJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultPrefetchPreload]
type zoneSettingsCollectionResultPrefetchPreloadJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultPrefetchPreload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultPrefetchPreload) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultPrefetchPreloadID string

const (
	ZoneSettingsCollectionResultPrefetchPreloadIDPrefetchPreload ZoneSettingsCollectionResultPrefetchPreloadID = "prefetch_preload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultPrefetchPreloadEditable bool

const (
	ZoneSettingsCollectionResultPrefetchPreloadEditableTrue  ZoneSettingsCollectionResultPrefetchPreloadEditable = true
	ZoneSettingsCollectionResultPrefetchPreloadEditableFalse ZoneSettingsCollectionResultPrefetchPreloadEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultPrefetchPreloadValue string

const (
	ZoneSettingsCollectionResultPrefetchPreloadValueOn  ZoneSettingsCollectionResultPrefetchPreloadValue = "on"
	ZoneSettingsCollectionResultPrefetchPreloadValueOff ZoneSettingsCollectionResultPrefetchPreloadValue = "off"
)

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
type ZoneSettingsCollectionResultPrivacyPass struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultPrivacyPassID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultPrivacyPassEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultPrivacyPassValue `json:"value"`
	JSON  zoneSettingsCollectionResultPrivacyPassJSON  `json:"-"`
}

// zoneSettingsCollectionResultPrivacyPassJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultPrivacyPass]
type zoneSettingsCollectionResultPrivacyPassJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultPrivacyPass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultPrivacyPass) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultPrivacyPassID string

const (
	ZoneSettingsCollectionResultPrivacyPassIDPrivacyPass ZoneSettingsCollectionResultPrivacyPassID = "privacy_pass"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultPrivacyPassEditable bool

const (
	ZoneSettingsCollectionResultPrivacyPassEditableTrue  ZoneSettingsCollectionResultPrivacyPassEditable = true
	ZoneSettingsCollectionResultPrivacyPassEditableFalse ZoneSettingsCollectionResultPrivacyPassEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultPrivacyPassValue string

const (
	ZoneSettingsCollectionResultPrivacyPassValueOn  ZoneSettingsCollectionResultPrivacyPassValue = "on"
	ZoneSettingsCollectionResultPrivacyPassValueOff ZoneSettingsCollectionResultPrivacyPassValue = "off"
)

// Maximum time between two read operations from origin.
type ZoneSettingsCollectionResultProxyReadTimeout struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultProxyReadTimeoutID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultProxyReadTimeoutEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value float64                                          `json:"value"`
	JSON  zoneSettingsCollectionResultProxyReadTimeoutJSON `json:"-"`
}

// zoneSettingsCollectionResultProxyReadTimeoutJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultProxyReadTimeout]
type zoneSettingsCollectionResultProxyReadTimeoutJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultProxyReadTimeout) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultProxyReadTimeout) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultProxyReadTimeoutID string

const (
	ZoneSettingsCollectionResultProxyReadTimeoutIDProxyReadTimeout ZoneSettingsCollectionResultProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultProxyReadTimeoutEditable bool

const (
	ZoneSettingsCollectionResultProxyReadTimeoutEditableTrue  ZoneSettingsCollectionResultProxyReadTimeoutEditable = true
	ZoneSettingsCollectionResultProxyReadTimeoutEditableFalse ZoneSettingsCollectionResultProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingsCollectionResultPseudoIpv4 struct {
	// Value of the Pseudo IPv4 setting.
	ID ZoneSettingsCollectionResultPseudoIpv4ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultPseudoIpv4Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the Pseudo IPv4 setting.
	Value ZoneSettingsCollectionResultPseudoIpv4Value `json:"value"`
	JSON  zoneSettingsCollectionResultPseudoIpv4JSON  `json:"-"`
}

// zoneSettingsCollectionResultPseudoIpv4JSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultPseudoIpv4]
type zoneSettingsCollectionResultPseudoIpv4JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultPseudoIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultPseudoIpv4) implementsZoneSettingsCollectionResult() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingsCollectionResultPseudoIpv4ID string

const (
	ZoneSettingsCollectionResultPseudoIpv4IDPseudoIpv4 ZoneSettingsCollectionResultPseudoIpv4ID = "pseudo_ipv4"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultPseudoIpv4Editable bool

const (
	ZoneSettingsCollectionResultPseudoIpv4EditableTrue  ZoneSettingsCollectionResultPseudoIpv4Editable = true
	ZoneSettingsCollectionResultPseudoIpv4EditableFalse ZoneSettingsCollectionResultPseudoIpv4Editable = false
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingsCollectionResultPseudoIpv4Value string

const (
	ZoneSettingsCollectionResultPseudoIpv4ValueOff             ZoneSettingsCollectionResultPseudoIpv4Value = "off"
	ZoneSettingsCollectionResultPseudoIpv4ValueAddHeader       ZoneSettingsCollectionResultPseudoIpv4Value = "add_header"
	ZoneSettingsCollectionResultPseudoIpv4ValueOverwriteHeader ZoneSettingsCollectionResultPseudoIpv4Value = "overwrite_header"
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingsCollectionResultResponseBuffering struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultResponseBufferingID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultResponseBufferingEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultResponseBufferingValue `json:"value"`
	JSON  zoneSettingsCollectionResultResponseBufferingJSON  `json:"-"`
}

// zoneSettingsCollectionResultResponseBufferingJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultResponseBuffering]
type zoneSettingsCollectionResultResponseBufferingJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultResponseBuffering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultResponseBuffering) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultResponseBufferingID string

const (
	ZoneSettingsCollectionResultResponseBufferingIDResponseBuffering ZoneSettingsCollectionResultResponseBufferingID = "response_buffering"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultResponseBufferingEditable bool

const (
	ZoneSettingsCollectionResultResponseBufferingEditableTrue  ZoneSettingsCollectionResultResponseBufferingEditable = true
	ZoneSettingsCollectionResultResponseBufferingEditableFalse ZoneSettingsCollectionResultResponseBufferingEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultResponseBufferingValue string

const (
	ZoneSettingsCollectionResultResponseBufferingValueOn  ZoneSettingsCollectionResultResponseBufferingValue = "on"
	ZoneSettingsCollectionResultResponseBufferingValueOff ZoneSettingsCollectionResultResponseBufferingValue = "off"
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
type ZoneSettingsCollectionResultRocketLoader struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultRocketLoaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultRocketLoaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultRocketLoaderValue `json:"value"`
	JSON  zoneSettingsCollectionResultRocketLoaderJSON  `json:"-"`
}

// zoneSettingsCollectionResultRocketLoaderJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultRocketLoader]
type zoneSettingsCollectionResultRocketLoaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultRocketLoader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultRocketLoader) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultRocketLoaderID string

const (
	ZoneSettingsCollectionResultRocketLoaderIDRocketLoader ZoneSettingsCollectionResultRocketLoaderID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultRocketLoaderEditable bool

const (
	ZoneSettingsCollectionResultRocketLoaderEditableTrue  ZoneSettingsCollectionResultRocketLoaderEditable = true
	ZoneSettingsCollectionResultRocketLoaderEditableFalse ZoneSettingsCollectionResultRocketLoaderEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultRocketLoaderValue string

const (
	ZoneSettingsCollectionResultRocketLoaderValueOn  ZoneSettingsCollectionResultRocketLoaderValue = "on"
	ZoneSettingsCollectionResultRocketLoaderValueOff ZoneSettingsCollectionResultRocketLoaderValue = "off"
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                                             `json:"modified_on,nullable" format:"date-time"`
	Value      ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValue `json:"value"`
	JSON       zoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationJSON  `json:"-"`
}

// zoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationJSON contains
// the JSON metadata for the struct
// [ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimization]
type zoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimization) implementsZoneSettingsCollectionResult() {
}

// ID of the zone setting.
type ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationEditable = false
)

type ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValue struct {
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
	JSON     zoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValueJSON `json:"-"`
}

// zoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValueJSON
// contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValue]
type zoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValueJSON struct {
	CacheByDeviceType apijson.Field
	Cf                apijson.Field
	Enabled           apijson.Field
	Hostnames         apijson.Field
	Wordpress         apijson.Field
	WpPlugin          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSchemasAutomaticPlatformOptimizationValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Cloudflare security header for a zone.
type ZoneSettingsCollectionResultSecurityHeader struct {
	// ID of the zone's security header.
	ID ZoneSettingsCollectionResultSecurityHeaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSecurityHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time                                       `json:"modified_on,nullable" format:"date-time"`
	Value      ZoneSettingsCollectionResultSecurityHeaderValue `json:"value"`
	JSON       zoneSettingsCollectionResultSecurityHeaderJSON  `json:"-"`
}

// zoneSettingsCollectionResultSecurityHeaderJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultSecurityHeader]
type zoneSettingsCollectionResultSecurityHeaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSecurityHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSecurityHeader) implementsZoneSettingsCollectionResult() {}

// ID of the zone's security header.
type ZoneSettingsCollectionResultSecurityHeaderID string

const (
	ZoneSettingsCollectionResultSecurityHeaderIDSecurityHeader ZoneSettingsCollectionResultSecurityHeaderID = "security_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSecurityHeaderEditable bool

const (
	ZoneSettingsCollectionResultSecurityHeaderEditableTrue  ZoneSettingsCollectionResultSecurityHeaderEditable = true
	ZoneSettingsCollectionResultSecurityHeaderEditableFalse ZoneSettingsCollectionResultSecurityHeaderEditable = false
)

type ZoneSettingsCollectionResultSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity ZoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurity `json:"strict_transport_security"`
	JSON                    zoneSettingsCollectionResultSecurityHeaderValueJSON                    `json:"-"`
}

// zoneSettingsCollectionResultSecurityHeaderValueJSON contains the JSON metadata
// for the struct [ZoneSettingsCollectionResultSecurityHeaderValue]
type zoneSettingsCollectionResultSecurityHeaderValueJSON struct {
	StrictTransportSecurity apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSecurityHeaderValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Strict Transport Security.
type ZoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled bool `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains bool `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge float64 `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff bool                                                                       `json:"nosniff"`
	JSON    zoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurityJSON `json:"-"`
}

// zoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurityJSON
// contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurity]
type zoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurityJSON struct {
	Enabled           apijson.Field
	IncludeSubdomains apijson.Field
	MaxAge            apijson.Field
	Nosniff           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSecurityHeaderValueStrictTransportSecurity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingsCollectionResultSecurityLevel struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultSecurityLevelID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSecurityLevelEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultSecurityLevelValue `json:"value"`
	JSON  zoneSettingsCollectionResultSecurityLevelJSON  `json:"-"`
}

// zoneSettingsCollectionResultSecurityLevelJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultSecurityLevel]
type zoneSettingsCollectionResultSecurityLevelJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSecurityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSecurityLevel) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultSecurityLevelID string

const (
	ZoneSettingsCollectionResultSecurityLevelIDSecurityLevel ZoneSettingsCollectionResultSecurityLevelID = "security_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSecurityLevelEditable bool

const (
	ZoneSettingsCollectionResultSecurityLevelEditableTrue  ZoneSettingsCollectionResultSecurityLevelEditable = true
	ZoneSettingsCollectionResultSecurityLevelEditableFalse ZoneSettingsCollectionResultSecurityLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultSecurityLevelValue string

const (
	ZoneSettingsCollectionResultSecurityLevelValueOff            ZoneSettingsCollectionResultSecurityLevelValue = "off"
	ZoneSettingsCollectionResultSecurityLevelValueEssentiallyOff ZoneSettingsCollectionResultSecurityLevelValue = "essentially_off"
	ZoneSettingsCollectionResultSecurityLevelValueLow            ZoneSettingsCollectionResultSecurityLevelValue = "low"
	ZoneSettingsCollectionResultSecurityLevelValueMedium         ZoneSettingsCollectionResultSecurityLevelValue = "medium"
	ZoneSettingsCollectionResultSecurityLevelValueHigh           ZoneSettingsCollectionResultSecurityLevelValue = "high"
	ZoneSettingsCollectionResultSecurityLevelValueUnderAttack    ZoneSettingsCollectionResultSecurityLevelValue = "under_attack"
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
type ZoneSettingsCollectionResultServerSideExclude struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultServerSideExcludeID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultServerSideExcludeEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultServerSideExcludeValue `json:"value"`
	JSON  zoneSettingsCollectionResultServerSideExcludeJSON  `json:"-"`
}

// zoneSettingsCollectionResultServerSideExcludeJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultServerSideExclude]
type zoneSettingsCollectionResultServerSideExcludeJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultServerSideExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultServerSideExclude) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultServerSideExcludeID string

const (
	ZoneSettingsCollectionResultServerSideExcludeIDServerSideExclude ZoneSettingsCollectionResultServerSideExcludeID = "server_side_exclude"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultServerSideExcludeEditable bool

const (
	ZoneSettingsCollectionResultServerSideExcludeEditableTrue  ZoneSettingsCollectionResultServerSideExcludeEditable = true
	ZoneSettingsCollectionResultServerSideExcludeEditableFalse ZoneSettingsCollectionResultServerSideExcludeEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultServerSideExcludeValue string

const (
	ZoneSettingsCollectionResultServerSideExcludeValueOn  ZoneSettingsCollectionResultServerSideExcludeValue = "on"
	ZoneSettingsCollectionResultServerSideExcludeValueOff ZoneSettingsCollectionResultServerSideExcludeValue = "off"
)

// Allow SHA1 support.
type ZoneSettingsCollectionResultSha1Support struct {
	// Zone setting identifier.
	ID ZoneSettingsCollectionResultSha1SupportID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSha1SupportEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultSha1SupportValue `json:"value"`
	JSON  zoneSettingsCollectionResultSha1SupportJSON  `json:"-"`
}

// zoneSettingsCollectionResultSha1SupportJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultSha1Support]
type zoneSettingsCollectionResultSha1SupportJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSha1Support) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSha1Support) implementsZoneSettingsCollectionResult() {}

// Zone setting identifier.
type ZoneSettingsCollectionResultSha1SupportID string

const (
	ZoneSettingsCollectionResultSha1SupportIDSha1Support ZoneSettingsCollectionResultSha1SupportID = "sha1_support"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSha1SupportEditable bool

const (
	ZoneSettingsCollectionResultSha1SupportEditableTrue  ZoneSettingsCollectionResultSha1SupportEditable = true
	ZoneSettingsCollectionResultSha1SupportEditableFalse ZoneSettingsCollectionResultSha1SupportEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultSha1SupportValue string

const (
	ZoneSettingsCollectionResultSha1SupportValueOff ZoneSettingsCollectionResultSha1SupportValue = "off"
	ZoneSettingsCollectionResultSha1SupportValueOn  ZoneSettingsCollectionResultSha1SupportValue = "on"
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingsCollectionResultSortQueryStringForCache struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultSortQueryStringForCacheID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSortQueryStringForCacheEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultSortQueryStringForCacheValue `json:"value"`
	JSON  zoneSettingsCollectionResultSortQueryStringForCacheJSON  `json:"-"`
}

// zoneSettingsCollectionResultSortQueryStringForCacheJSON contains the JSON
// metadata for the struct [ZoneSettingsCollectionResultSortQueryStringForCache]
type zoneSettingsCollectionResultSortQueryStringForCacheJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSortQueryStringForCache) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSortQueryStringForCache) implementsZoneSettingsCollectionResult() {
}

// ID of the zone setting.
type ZoneSettingsCollectionResultSortQueryStringForCacheID string

const (
	ZoneSettingsCollectionResultSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingsCollectionResultSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSortQueryStringForCacheEditable bool

const (
	ZoneSettingsCollectionResultSortQueryStringForCacheEditableTrue  ZoneSettingsCollectionResultSortQueryStringForCacheEditable = true
	ZoneSettingsCollectionResultSortQueryStringForCacheEditableFalse ZoneSettingsCollectionResultSortQueryStringForCacheEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultSortQueryStringForCacheValue string

const (
	ZoneSettingsCollectionResultSortQueryStringForCacheValueOn  ZoneSettingsCollectionResultSortQueryStringForCacheValue = "on"
	ZoneSettingsCollectionResultSortQueryStringForCacheValueOff ZoneSettingsCollectionResultSortQueryStringForCacheValue = "off"
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
type ZoneSettingsCollectionResultSsl struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultSslID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSslEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingsCollectionResultSslValue `json:"value"`
	JSON  zoneSettingsCollectionResultSslJSON  `json:"-"`
}

// zoneSettingsCollectionResultSslJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultSsl]
type zoneSettingsCollectionResultSslJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSsl) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultSslID string

const (
	ZoneSettingsCollectionResultSslIDSsl ZoneSettingsCollectionResultSslID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSslEditable bool

const (
	ZoneSettingsCollectionResultSslEditableTrue  ZoneSettingsCollectionResultSslEditable = true
	ZoneSettingsCollectionResultSslEditableFalse ZoneSettingsCollectionResultSslEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingsCollectionResultSslValue string

const (
	ZoneSettingsCollectionResultSslValueOff      ZoneSettingsCollectionResultSslValue = "off"
	ZoneSettingsCollectionResultSslValueFlexible ZoneSettingsCollectionResultSslValue = "flexible"
	ZoneSettingsCollectionResultSslValueFull     ZoneSettingsCollectionResultSslValue = "full"
	ZoneSettingsCollectionResultSslValueStrict   ZoneSettingsCollectionResultSslValue = "strict"
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingsCollectionResultSslRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingsCollectionResultSslRecommenderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultSslRecommenderEditable `json:"editable"`
	// ssl-recommender enrollment setting.
	Enabled bool `json:"enabled"`
	// last time this setting was modified.
	ModifiedOn time.Time                                      `json:"modified_on,nullable" format:"date-time"`
	JSON       zoneSettingsCollectionResultSslRecommenderJSON `json:"-"`
}

// zoneSettingsCollectionResultSslRecommenderJSON contains the JSON metadata for
// the struct [ZoneSettingsCollectionResultSslRecommender]
type zoneSettingsCollectionResultSslRecommenderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Enabled     apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultSslRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultSslRecommender) implementsZoneSettingsCollectionResult() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingsCollectionResultSslRecommenderID string

const (
	ZoneSettingsCollectionResultSslRecommenderIDSslRecommender ZoneSettingsCollectionResultSslRecommenderID = "ssl_recommender"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultSslRecommenderEditable bool

const (
	ZoneSettingsCollectionResultSslRecommenderEditableTrue  ZoneSettingsCollectionResultSslRecommenderEditable = true
	ZoneSettingsCollectionResultSslRecommenderEditableFalse ZoneSettingsCollectionResultSslRecommenderEditable = false
)

// Only allows TLS1.2.
type ZoneSettingsCollectionResultTls1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingsCollectionResultTls1_2OnlyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultTls1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultTls1_2OnlyValue `json:"value"`
	JSON  zoneSettingsCollectionResultTls1_2OnlyJSON  `json:"-"`
}

// zoneSettingsCollectionResultTls1_2OnlyJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultTls1_2Only]
type zoneSettingsCollectionResultTls1_2OnlyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultTls1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultTls1_2Only) implementsZoneSettingsCollectionResult() {}

// Zone setting identifier.
type ZoneSettingsCollectionResultTls1_2OnlyID string

const (
	ZoneSettingsCollectionResultTls1_2OnlyIDTls1_2Only ZoneSettingsCollectionResultTls1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultTls1_2OnlyEditable bool

const (
	ZoneSettingsCollectionResultTls1_2OnlyEditableTrue  ZoneSettingsCollectionResultTls1_2OnlyEditable = true
	ZoneSettingsCollectionResultTls1_2OnlyEditableFalse ZoneSettingsCollectionResultTls1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultTls1_2OnlyValue string

const (
	ZoneSettingsCollectionResultTls1_2OnlyValueOff ZoneSettingsCollectionResultTls1_2OnlyValue = "off"
	ZoneSettingsCollectionResultTls1_2OnlyValueOn  ZoneSettingsCollectionResultTls1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingsCollectionResultTls1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultTls1_3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultTls1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingsCollectionResultTls1_3Value `json:"value"`
	JSON  zoneSettingsCollectionResultTls1_3JSON  `json:"-"`
}

// zoneSettingsCollectionResultTls1_3JSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultTls1_3]
type zoneSettingsCollectionResultTls1_3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultTls1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultTls1_3) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultTls1_3ID string

const (
	ZoneSettingsCollectionResultTls1_3IDTls1_3 ZoneSettingsCollectionResultTls1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultTls1_3Editable bool

const (
	ZoneSettingsCollectionResultTls1_3EditableTrue  ZoneSettingsCollectionResultTls1_3Editable = true
	ZoneSettingsCollectionResultTls1_3EditableFalse ZoneSettingsCollectionResultTls1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingsCollectionResultTls1_3Value string

const (
	ZoneSettingsCollectionResultTls1_3ValueOn  ZoneSettingsCollectionResultTls1_3Value = "on"
	ZoneSettingsCollectionResultTls1_3ValueOff ZoneSettingsCollectionResultTls1_3Value = "off"
	ZoneSettingsCollectionResultTls1_3ValueZrt ZoneSettingsCollectionResultTls1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingsCollectionResultTlsClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultTlsClientAuthID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultTlsClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingsCollectionResultTlsClientAuthValue `json:"value"`
	JSON  zoneSettingsCollectionResultTlsClientAuthJSON  `json:"-"`
}

// zoneSettingsCollectionResultTlsClientAuthJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultTlsClientAuth]
type zoneSettingsCollectionResultTlsClientAuthJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultTlsClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultTlsClientAuth) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultTlsClientAuthID string

const (
	ZoneSettingsCollectionResultTlsClientAuthIDTlsClientAuth ZoneSettingsCollectionResultTlsClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultTlsClientAuthEditable bool

const (
	ZoneSettingsCollectionResultTlsClientAuthEditableTrue  ZoneSettingsCollectionResultTlsClientAuthEditable = true
	ZoneSettingsCollectionResultTlsClientAuthEditableFalse ZoneSettingsCollectionResultTlsClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingsCollectionResultTlsClientAuthValue string

const (
	ZoneSettingsCollectionResultTlsClientAuthValueOn  ZoneSettingsCollectionResultTlsClientAuthValue = "on"
	ZoneSettingsCollectionResultTlsClientAuthValueOff ZoneSettingsCollectionResultTlsClientAuthValue = "off"
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingsCollectionResultTrueClientIPHeader struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultTrueClientIPHeaderID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultTrueClientIPHeaderEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultTrueClientIPHeaderValue `json:"value"`
	JSON  zoneSettingsCollectionResultTrueClientIPHeaderJSON  `json:"-"`
}

// zoneSettingsCollectionResultTrueClientIPHeaderJSON contains the JSON metadata
// for the struct [ZoneSettingsCollectionResultTrueClientIPHeader]
type zoneSettingsCollectionResultTrueClientIPHeaderJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultTrueClientIPHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultTrueClientIPHeader) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultTrueClientIPHeaderID string

const (
	ZoneSettingsCollectionResultTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingsCollectionResultTrueClientIPHeaderID = "true_client_ip_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultTrueClientIPHeaderEditable bool

const (
	ZoneSettingsCollectionResultTrueClientIPHeaderEditableTrue  ZoneSettingsCollectionResultTrueClientIPHeaderEditable = true
	ZoneSettingsCollectionResultTrueClientIPHeaderEditableFalse ZoneSettingsCollectionResultTrueClientIPHeaderEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultTrueClientIPHeaderValue string

const (
	ZoneSettingsCollectionResultTrueClientIPHeaderValueOn  ZoneSettingsCollectionResultTrueClientIPHeaderValue = "on"
	ZoneSettingsCollectionResultTrueClientIPHeaderValueOff ZoneSettingsCollectionResultTrueClientIPHeaderValue = "off"
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
type ZoneSettingsCollectionResultWaf struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultWafID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultWafEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultWafValue `json:"value"`
	JSON  zoneSettingsCollectionResultWafJSON  `json:"-"`
}

// zoneSettingsCollectionResultWafJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultWaf]
type zoneSettingsCollectionResultWafJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultWaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultWaf) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultWafID string

const (
	ZoneSettingsCollectionResultWafIDWaf ZoneSettingsCollectionResultWafID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultWafEditable bool

const (
	ZoneSettingsCollectionResultWafEditableTrue  ZoneSettingsCollectionResultWafEditable = true
	ZoneSettingsCollectionResultWafEditableFalse ZoneSettingsCollectionResultWafEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultWafValue string

const (
	ZoneSettingsCollectionResultWafValueOn  ZoneSettingsCollectionResultWafValue = "on"
	ZoneSettingsCollectionResultWafValueOff ZoneSettingsCollectionResultWafValue = "off"
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingsCollectionResultWebp struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultWebpID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultWebpEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultWebpValue `json:"value"`
	JSON  zoneSettingsCollectionResultWebpJSON  `json:"-"`
}

// zoneSettingsCollectionResultWebpJSON contains the JSON metadata for the struct
// [ZoneSettingsCollectionResultWebp]
type zoneSettingsCollectionResultWebpJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultWebp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultWebp) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultWebpID string

const (
	ZoneSettingsCollectionResultWebpIDWebp ZoneSettingsCollectionResultWebpID = "webp"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultWebpEditable bool

const (
	ZoneSettingsCollectionResultWebpEditableTrue  ZoneSettingsCollectionResultWebpEditable = true
	ZoneSettingsCollectionResultWebpEditableFalse ZoneSettingsCollectionResultWebpEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultWebpValue string

const (
	ZoneSettingsCollectionResultWebpValueOff ZoneSettingsCollectionResultWebpValue = "off"
	ZoneSettingsCollectionResultWebpValueOn  ZoneSettingsCollectionResultWebpValue = "on"
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingsCollectionResultWebsockets struct {
	// ID of the zone setting.
	ID ZoneSettingsCollectionResultWebsocketsID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingsCollectionResultWebsocketsEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingsCollectionResultWebsocketsValue `json:"value"`
	JSON  zoneSettingsCollectionResultWebsocketsJSON  `json:"-"`
}

// zoneSettingsCollectionResultWebsocketsJSON contains the JSON metadata for the
// struct [ZoneSettingsCollectionResultWebsockets]
type zoneSettingsCollectionResultWebsocketsJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingsCollectionResultWebsockets) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingsCollectionResultWebsockets) implementsZoneSettingsCollectionResult() {}

// ID of the zone setting.
type ZoneSettingsCollectionResultWebsocketsID string

const (
	ZoneSettingsCollectionResultWebsocketsIDWebsockets ZoneSettingsCollectionResultWebsocketsID = "websockets"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingsCollectionResultWebsocketsEditable bool

const (
	ZoneSettingsCollectionResultWebsocketsEditableTrue  ZoneSettingsCollectionResultWebsocketsEditable = true
	ZoneSettingsCollectionResultWebsocketsEditableFalse ZoneSettingsCollectionResultWebsocketsEditable = false
)

// Value of the zone setting.
type ZoneSettingsCollectionResultWebsocketsValue string

const (
	ZoneSettingsCollectionResultWebsocketsValueOff ZoneSettingsCollectionResultWebsocketsValue = "off"
	ZoneSettingsCollectionResultWebsocketsValueOn  ZoneSettingsCollectionResultWebsocketsValue = "on"
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
// Satisfied by [ZeroRttParam], [ZoneSettingEditParamsItemsAdvancedDdos],
// [ZoneSettingEditParamsItemsAlwaysOnline],
// [ZoneSettingEditParamsItemsAlwaysUseHTTPs],
// [ZoneSettingEditParamsItemsAutomaticHTTPsRewrites],
// [ZoneSettingEditParamsItemsBrotli], [ZoneSettingEditParamsItemsBrowserCacheTtl],
// [ZoneSettingEditParamsItemsBrowserCheck],
// [ZoneSettingEditParamsItemsCacheLevel],
// [ZoneSettingEditParamsItemsChallengeTtl], [ZoneSettingEditParamsItemsCiphers],
// [ZoneSettingEditParamsItemsCnameFlattening],
// [ZoneSettingEditParamsItemsDevelopmentMode],
// [ZoneSettingEditParamsItemsEarlyHints],
// [ZoneSettingEditParamsItemsEdgeCacheTtl],
// [ZoneSettingEditParamsItemsEmailObfuscation],
// [ZoneSettingEditParamsItemsH2Prioritization],
// [ZoneSettingEditParamsItemsHotlinkProtection],
// [ZoneSettingEditParamsItemsHttp2], [ZoneSettingEditParamsItemsHttp3],
// [ZoneSettingEditParamsItemsImageResizing],
// [ZoneSettingEditParamsItemsIPGeolocation], [ZoneSettingEditParamsItemsIpv6],
// [ZoneSettingEditParamsItemsMaxUpload],
// [ZoneSettingEditParamsItemsMinTlsVersion], [ZoneSettingEditParamsItemsMinify],
// [ZoneSettingEditParamsItemsMirage], [ZoneSettingEditParamsItemsMobileRedirect],
// [ZoneSettingEditParamsItemsNel],
// [ZoneSettingEditParamsItemsOpportunisticEncryption],
// [ZoneSettingEditParamsItemsOpportunisticOnion],
// [ZoneSettingEditParamsItemsOrangeToOrange],
// [ZoneSettingEditParamsItemsOriginErrorPagePassThru],
// [ZoneSettingEditParamsItemsOriginMaxHTTPVersion],
// [ZoneSettingEditParamsItemsPolish], [ZoneSettingEditParamsItemsPrefetchPreload],
// [ZoneSettingEditParamsItemsPrivacyPass],
// [ZoneSettingEditParamsItemsProxyReadTimeout],
// [ZoneSettingEditParamsItemsPseudoIpv4],
// [ZoneSettingEditParamsItemsResponseBuffering],
// [ZoneSettingEditParamsItemsRocketLoader],
// [ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimization],
// [ZoneSettingEditParamsItemsSecurityHeader],
// [ZoneSettingEditParamsItemsSecurityLevel],
// [ZoneSettingEditParamsItemsServerSideExclude],
// [ZoneSettingEditParamsItemsSha1Support],
// [ZoneSettingEditParamsItemsSortQueryStringForCache],
// [ZoneSettingEditParamsItemsSsl], [ZoneSettingEditParamsItemsSslRecommender],
// [ZoneSettingEditParamsItemsTls1_2Only], [ZoneSettingEditParamsItemsTls1_3],
// [ZoneSettingEditParamsItemsTlsClientAuth],
// [ZoneSettingEditParamsItemsTrueClientIPHeader], [ZoneSettingEditParamsItemsWaf],
// [ZoneSettingEditParamsItemsWebp], [ZoneSettingEditParamsItemsWebsockets].
type ZoneSettingEditParamsItem interface {
	implementsZoneSettingEditParamsItem()
}

// Advanced protection from Distributed Denial of Service (DDoS) attacks on your
// website. This is an uneditable value that is 'on' in the case of Business and
// Enterprise zones.
type ZoneSettingEditParamsItemsAdvancedDdos struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsAdvancedDdosID] `json:"id"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value param.Field[ZoneSettingEditParamsItemsAdvancedDdosValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsAdvancedDdos) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsAdvancedDdos) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsAdvancedDdosID string

const (
	ZoneSettingEditParamsItemsAdvancedDdosIDAdvancedDdos ZoneSettingEditParamsItemsAdvancedDdosID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsAdvancedDdosEditable bool

const (
	ZoneSettingEditParamsItemsAdvancedDdosEditableTrue  ZoneSettingEditParamsItemsAdvancedDdosEditable = true
	ZoneSettingEditParamsItemsAdvancedDdosEditableFalse ZoneSettingEditParamsItemsAdvancedDdosEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingEditParamsItemsAdvancedDdosValue string

const (
	ZoneSettingEditParamsItemsAdvancedDdosValueOn  ZoneSettingEditParamsItemsAdvancedDdosValue = "on"
	ZoneSettingEditParamsItemsAdvancedDdosValueOff ZoneSettingEditParamsItemsAdvancedDdosValue = "off"
)

// When enabled, Cloudflare serves limited copies of web pages available from the
// [Internet Archive's Wayback Machine](https://archive.org/web/) if your server is
// offline. Refer to
// [Always Online](https://developers.cloudflare.com/cache/about/always-online) for
// more information.
type ZoneSettingEditParamsItemsAlwaysOnline struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsAlwaysOnlineID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsAlwaysOnlineValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsAlwaysOnline) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsAlwaysOnline) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsAlwaysOnlineID string

const (
	ZoneSettingEditParamsItemsAlwaysOnlineIDAlwaysOnline ZoneSettingEditParamsItemsAlwaysOnlineID = "always_online"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsAlwaysOnlineEditable bool

const (
	ZoneSettingEditParamsItemsAlwaysOnlineEditableTrue  ZoneSettingEditParamsItemsAlwaysOnlineEditable = true
	ZoneSettingEditParamsItemsAlwaysOnlineEditableFalse ZoneSettingEditParamsItemsAlwaysOnlineEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsAlwaysOnlineValue string

const (
	ZoneSettingEditParamsItemsAlwaysOnlineValueOn  ZoneSettingEditParamsItemsAlwaysOnlineValue = "on"
	ZoneSettingEditParamsItemsAlwaysOnlineValueOff ZoneSettingEditParamsItemsAlwaysOnlineValue = "off"
)

// Reply to all requests for URLs that use "http" with a 301 redirect to the
// equivalent "https" URL. If you only want to redirect for a subset of requests,
// consider creating an "Always use HTTPS" page rule.
type ZoneSettingEditParamsItemsAlwaysUseHTTPs struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsAlwaysUseHTTPsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsAlwaysUseHTTPsValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsAlwaysUseHTTPs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsAlwaysUseHTTPs) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsAlwaysUseHTTPsID string

const (
	ZoneSettingEditParamsItemsAlwaysUseHTTPsIDAlwaysUseHTTPs ZoneSettingEditParamsItemsAlwaysUseHTTPsID = "always_use_https"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsAlwaysUseHTTPsEditable bool

const (
	ZoneSettingEditParamsItemsAlwaysUseHTTPsEditableTrue  ZoneSettingEditParamsItemsAlwaysUseHTTPsEditable = true
	ZoneSettingEditParamsItemsAlwaysUseHTTPsEditableFalse ZoneSettingEditParamsItemsAlwaysUseHTTPsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsAlwaysUseHTTPsValue string

const (
	ZoneSettingEditParamsItemsAlwaysUseHTTPsValueOn  ZoneSettingEditParamsItemsAlwaysUseHTTPsValue = "on"
	ZoneSettingEditParamsItemsAlwaysUseHTTPsValueOff ZoneSettingEditParamsItemsAlwaysUseHTTPsValue = "off"
)

// Enable the Automatic HTTPS Rewrites feature for this zone.
type ZoneSettingEditParamsItemsAutomaticHTTPsRewrites struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsAutomaticHTTPsRewritesID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsAutomaticHTTPsRewritesValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsAutomaticHTTPsRewrites) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsAutomaticHTTPsRewrites) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsAutomaticHTTPsRewritesID string

const (
	ZoneSettingEditParamsItemsAutomaticHTTPsRewritesIDAutomaticHTTPsRewrites ZoneSettingEditParamsItemsAutomaticHTTPsRewritesID = "automatic_https_rewrites"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsAutomaticHTTPsRewritesEditable bool

const (
	ZoneSettingEditParamsItemsAutomaticHTTPsRewritesEditableTrue  ZoneSettingEditParamsItemsAutomaticHTTPsRewritesEditable = true
	ZoneSettingEditParamsItemsAutomaticHTTPsRewritesEditableFalse ZoneSettingEditParamsItemsAutomaticHTTPsRewritesEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsAutomaticHTTPsRewritesValue string

const (
	ZoneSettingEditParamsItemsAutomaticHTTPsRewritesValueOn  ZoneSettingEditParamsItemsAutomaticHTTPsRewritesValue = "on"
	ZoneSettingEditParamsItemsAutomaticHTTPsRewritesValueOff ZoneSettingEditParamsItemsAutomaticHTTPsRewritesValue = "off"
)

// When the client requesting an asset supports the Brotli compression algorithm,
// Cloudflare will serve a Brotli compressed version of the asset.
type ZoneSettingEditParamsItemsBrotli struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsBrotliID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsBrotliValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsBrotli) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsBrotli) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsBrotliID string

const (
	ZoneSettingEditParamsItemsBrotliIDBrotli ZoneSettingEditParamsItemsBrotliID = "brotli"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsBrotliEditable bool

const (
	ZoneSettingEditParamsItemsBrotliEditableTrue  ZoneSettingEditParamsItemsBrotliEditable = true
	ZoneSettingEditParamsItemsBrotliEditableFalse ZoneSettingEditParamsItemsBrotliEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsBrotliValue string

const (
	ZoneSettingEditParamsItemsBrotliValueOff ZoneSettingEditParamsItemsBrotliValue = "off"
	ZoneSettingEditParamsItemsBrotliValueOn  ZoneSettingEditParamsItemsBrotliValue = "on"
)

// Browser Cache TTL (in seconds) specifies how long Cloudflare-cached resources
// will remain on your visitors' computers. Cloudflare will honor any larger times
// specified by your server.
// (https://support.cloudflare.com/hc/en-us/articles/200168276).
type ZoneSettingEditParamsItemsBrowserCacheTtl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsBrowserCacheTtlID] `json:"id"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingEditParamsItemsBrowserCacheTtlValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsBrowserCacheTtl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsBrowserCacheTtl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsBrowserCacheTtlID string

const (
	ZoneSettingEditParamsItemsBrowserCacheTtlIDBrowserCacheTtl ZoneSettingEditParamsItemsBrowserCacheTtlID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsBrowserCacheTtlEditable bool

const (
	ZoneSettingEditParamsItemsBrowserCacheTtlEditableTrue  ZoneSettingEditParamsItemsBrowserCacheTtlEditable = true
	ZoneSettingEditParamsItemsBrowserCacheTtlEditableFalse ZoneSettingEditParamsItemsBrowserCacheTtlEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingEditParamsItemsBrowserCacheTtlValue float64

const (
	ZoneSettingEditParamsItemsBrowserCacheTtlValue0        ZoneSettingEditParamsItemsBrowserCacheTtlValue = 0
	ZoneSettingEditParamsItemsBrowserCacheTtlValue30       ZoneSettingEditParamsItemsBrowserCacheTtlValue = 30
	ZoneSettingEditParamsItemsBrowserCacheTtlValue60       ZoneSettingEditParamsItemsBrowserCacheTtlValue = 60
	ZoneSettingEditParamsItemsBrowserCacheTtlValue120      ZoneSettingEditParamsItemsBrowserCacheTtlValue = 120
	ZoneSettingEditParamsItemsBrowserCacheTtlValue300      ZoneSettingEditParamsItemsBrowserCacheTtlValue = 300
	ZoneSettingEditParamsItemsBrowserCacheTtlValue1200     ZoneSettingEditParamsItemsBrowserCacheTtlValue = 1200
	ZoneSettingEditParamsItemsBrowserCacheTtlValue1800     ZoneSettingEditParamsItemsBrowserCacheTtlValue = 1800
	ZoneSettingEditParamsItemsBrowserCacheTtlValue3600     ZoneSettingEditParamsItemsBrowserCacheTtlValue = 3600
	ZoneSettingEditParamsItemsBrowserCacheTtlValue7200     ZoneSettingEditParamsItemsBrowserCacheTtlValue = 7200
	ZoneSettingEditParamsItemsBrowserCacheTtlValue10800    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 10800
	ZoneSettingEditParamsItemsBrowserCacheTtlValue14400    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 14400
	ZoneSettingEditParamsItemsBrowserCacheTtlValue18000    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 18000
	ZoneSettingEditParamsItemsBrowserCacheTtlValue28800    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 28800
	ZoneSettingEditParamsItemsBrowserCacheTtlValue43200    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 43200
	ZoneSettingEditParamsItemsBrowserCacheTtlValue57600    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 57600
	ZoneSettingEditParamsItemsBrowserCacheTtlValue72000    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 72000
	ZoneSettingEditParamsItemsBrowserCacheTtlValue86400    ZoneSettingEditParamsItemsBrowserCacheTtlValue = 86400
	ZoneSettingEditParamsItemsBrowserCacheTtlValue172800   ZoneSettingEditParamsItemsBrowserCacheTtlValue = 172800
	ZoneSettingEditParamsItemsBrowserCacheTtlValue259200   ZoneSettingEditParamsItemsBrowserCacheTtlValue = 259200
	ZoneSettingEditParamsItemsBrowserCacheTtlValue345600   ZoneSettingEditParamsItemsBrowserCacheTtlValue = 345600
	ZoneSettingEditParamsItemsBrowserCacheTtlValue432000   ZoneSettingEditParamsItemsBrowserCacheTtlValue = 432000
	ZoneSettingEditParamsItemsBrowserCacheTtlValue691200   ZoneSettingEditParamsItemsBrowserCacheTtlValue = 691200
	ZoneSettingEditParamsItemsBrowserCacheTtlValue1382400  ZoneSettingEditParamsItemsBrowserCacheTtlValue = 1382400
	ZoneSettingEditParamsItemsBrowserCacheTtlValue2073600  ZoneSettingEditParamsItemsBrowserCacheTtlValue = 2073600
	ZoneSettingEditParamsItemsBrowserCacheTtlValue2678400  ZoneSettingEditParamsItemsBrowserCacheTtlValue = 2678400
	ZoneSettingEditParamsItemsBrowserCacheTtlValue5356800  ZoneSettingEditParamsItemsBrowserCacheTtlValue = 5356800
	ZoneSettingEditParamsItemsBrowserCacheTtlValue16070400 ZoneSettingEditParamsItemsBrowserCacheTtlValue = 16070400
	ZoneSettingEditParamsItemsBrowserCacheTtlValue31536000 ZoneSettingEditParamsItemsBrowserCacheTtlValue = 31536000
)

// Browser Integrity Check is similar to Bad Behavior and looks for common HTTP
// headers abused most commonly by spammers and denies access to your page. It will
// also challenge visitors that do not have a user agent or a non standard user
// agent (also commonly used by abuse bots, crawlers or visitors).
// (https://support.cloudflare.com/hc/en-us/articles/200170086).
type ZoneSettingEditParamsItemsBrowserCheck struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsBrowserCheckID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsBrowserCheckValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsBrowserCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsBrowserCheck) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsBrowserCheckID string

const (
	ZoneSettingEditParamsItemsBrowserCheckIDBrowserCheck ZoneSettingEditParamsItemsBrowserCheckID = "browser_check"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsBrowserCheckEditable bool

const (
	ZoneSettingEditParamsItemsBrowserCheckEditableTrue  ZoneSettingEditParamsItemsBrowserCheckEditable = true
	ZoneSettingEditParamsItemsBrowserCheckEditableFalse ZoneSettingEditParamsItemsBrowserCheckEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsBrowserCheckValue string

const (
	ZoneSettingEditParamsItemsBrowserCheckValueOn  ZoneSettingEditParamsItemsBrowserCheckValue = "on"
	ZoneSettingEditParamsItemsBrowserCheckValueOff ZoneSettingEditParamsItemsBrowserCheckValue = "off"
)

// Cache Level functions based off the setting level. The basic setting will cache
// most static resources (i.e., css, images, and JavaScript). The simplified
// setting will ignore the query string when delivering a cached resource. The
// aggressive setting will cache all static resources, including ones with a query
// string. (https://support.cloudflare.com/hc/en-us/articles/200168256).
type ZoneSettingEditParamsItemsCacheLevel struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsCacheLevelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsCacheLevelValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsCacheLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsCacheLevel) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsCacheLevelID string

const (
	ZoneSettingEditParamsItemsCacheLevelIDCacheLevel ZoneSettingEditParamsItemsCacheLevelID = "cache_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsCacheLevelEditable bool

const (
	ZoneSettingEditParamsItemsCacheLevelEditableTrue  ZoneSettingEditParamsItemsCacheLevelEditable = true
	ZoneSettingEditParamsItemsCacheLevelEditableFalse ZoneSettingEditParamsItemsCacheLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsCacheLevelValue string

const (
	ZoneSettingEditParamsItemsCacheLevelValueAggressive ZoneSettingEditParamsItemsCacheLevelValue = "aggressive"
	ZoneSettingEditParamsItemsCacheLevelValueBasic      ZoneSettingEditParamsItemsCacheLevelValue = "basic"
	ZoneSettingEditParamsItemsCacheLevelValueSimplified ZoneSettingEditParamsItemsCacheLevelValue = "simplified"
)

// Specify how long a visitor is allowed access to your site after successfully
// completing a challenge (such as a CAPTCHA). After the TTL has expired the
// visitor will have to complete a new challenge. We recommend a 15 - 45 minute
// setting and will attempt to honor any setting above 45 minutes.
// (https://support.cloudflare.com/hc/en-us/articles/200170136).
type ZoneSettingEditParamsItemsChallengeTtl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsChallengeTtlID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsChallengeTtlValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsChallengeTtl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsChallengeTtl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsChallengeTtlID string

const (
	ZoneSettingEditParamsItemsChallengeTtlIDChallengeTtl ZoneSettingEditParamsItemsChallengeTtlID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsChallengeTtlEditable bool

const (
	ZoneSettingEditParamsItemsChallengeTtlEditableTrue  ZoneSettingEditParamsItemsChallengeTtlEditable = true
	ZoneSettingEditParamsItemsChallengeTtlEditableFalse ZoneSettingEditParamsItemsChallengeTtlEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsChallengeTtlValue float64

const (
	ZoneSettingEditParamsItemsChallengeTtlValue300      ZoneSettingEditParamsItemsChallengeTtlValue = 300
	ZoneSettingEditParamsItemsChallengeTtlValue900      ZoneSettingEditParamsItemsChallengeTtlValue = 900
	ZoneSettingEditParamsItemsChallengeTtlValue1800     ZoneSettingEditParamsItemsChallengeTtlValue = 1800
	ZoneSettingEditParamsItemsChallengeTtlValue2700     ZoneSettingEditParamsItemsChallengeTtlValue = 2700
	ZoneSettingEditParamsItemsChallengeTtlValue3600     ZoneSettingEditParamsItemsChallengeTtlValue = 3600
	ZoneSettingEditParamsItemsChallengeTtlValue7200     ZoneSettingEditParamsItemsChallengeTtlValue = 7200
	ZoneSettingEditParamsItemsChallengeTtlValue10800    ZoneSettingEditParamsItemsChallengeTtlValue = 10800
	ZoneSettingEditParamsItemsChallengeTtlValue14400    ZoneSettingEditParamsItemsChallengeTtlValue = 14400
	ZoneSettingEditParamsItemsChallengeTtlValue28800    ZoneSettingEditParamsItemsChallengeTtlValue = 28800
	ZoneSettingEditParamsItemsChallengeTtlValue57600    ZoneSettingEditParamsItemsChallengeTtlValue = 57600
	ZoneSettingEditParamsItemsChallengeTtlValue86400    ZoneSettingEditParamsItemsChallengeTtlValue = 86400
	ZoneSettingEditParamsItemsChallengeTtlValue604800   ZoneSettingEditParamsItemsChallengeTtlValue = 604800
	ZoneSettingEditParamsItemsChallengeTtlValue2592000  ZoneSettingEditParamsItemsChallengeTtlValue = 2592000
	ZoneSettingEditParamsItemsChallengeTtlValue31536000 ZoneSettingEditParamsItemsChallengeTtlValue = 31536000
)

// An allowlist of ciphers for TLS termination. These ciphers must be in the
// BoringSSL format.
type ZoneSettingEditParamsItemsCiphers struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsCiphersID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[[]string] `json:"value"`
}

func (r ZoneSettingEditParamsItemsCiphers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsCiphers) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsCiphersID string

const (
	ZoneSettingEditParamsItemsCiphersIDCiphers ZoneSettingEditParamsItemsCiphersID = "ciphers"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsCiphersEditable bool

const (
	ZoneSettingEditParamsItemsCiphersEditableTrue  ZoneSettingEditParamsItemsCiphersEditable = true
	ZoneSettingEditParamsItemsCiphersEditableFalse ZoneSettingEditParamsItemsCiphersEditable = false
)

// Whether or not cname flattening is on.
type ZoneSettingEditParamsItemsCnameFlattening struct {
	// How to flatten the cname destination.
	ID param.Field[ZoneSettingEditParamsItemsCnameFlatteningID] `json:"id"`
	// Value of the cname flattening setting.
	Value param.Field[ZoneSettingEditParamsItemsCnameFlatteningValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsCnameFlattening) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsCnameFlattening) implementsZoneSettingEditParamsItem() {}

// How to flatten the cname destination.
type ZoneSettingEditParamsItemsCnameFlatteningID string

const (
	ZoneSettingEditParamsItemsCnameFlatteningIDCnameFlattening ZoneSettingEditParamsItemsCnameFlatteningID = "cname_flattening"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsCnameFlatteningEditable bool

const (
	ZoneSettingEditParamsItemsCnameFlatteningEditableTrue  ZoneSettingEditParamsItemsCnameFlatteningEditable = true
	ZoneSettingEditParamsItemsCnameFlatteningEditableFalse ZoneSettingEditParamsItemsCnameFlatteningEditable = false
)

// Value of the cname flattening setting.
type ZoneSettingEditParamsItemsCnameFlatteningValue string

const (
	ZoneSettingEditParamsItemsCnameFlatteningValueFlattenAtRoot ZoneSettingEditParamsItemsCnameFlatteningValue = "flatten_at_root"
	ZoneSettingEditParamsItemsCnameFlatteningValueFlattenAll    ZoneSettingEditParamsItemsCnameFlatteningValue = "flatten_all"
)

// Development Mode temporarily allows you to enter development mode for your
// websites if you need to make changes to your site. This will bypass Cloudflare's
// accelerated cache and slow down your site, but is useful if you are making
// changes to cacheable content (like images, css, or JavaScript) and would like to
// see those changes right away. Once entered, development mode will last for 3
// hours and then automatically toggle off.
type ZoneSettingEditParamsItemsDevelopmentMode struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsDevelopmentModeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsDevelopmentModeValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsDevelopmentMode) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsDevelopmentMode) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsDevelopmentModeID string

const (
	ZoneSettingEditParamsItemsDevelopmentModeIDDevelopmentMode ZoneSettingEditParamsItemsDevelopmentModeID = "development_mode"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsDevelopmentModeEditable bool

const (
	ZoneSettingEditParamsItemsDevelopmentModeEditableTrue  ZoneSettingEditParamsItemsDevelopmentModeEditable = true
	ZoneSettingEditParamsItemsDevelopmentModeEditableFalse ZoneSettingEditParamsItemsDevelopmentModeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsDevelopmentModeValue string

const (
	ZoneSettingEditParamsItemsDevelopmentModeValueOn  ZoneSettingEditParamsItemsDevelopmentModeValue = "on"
	ZoneSettingEditParamsItemsDevelopmentModeValueOff ZoneSettingEditParamsItemsDevelopmentModeValue = "off"
)

// When enabled, Cloudflare will attempt to speed up overall page loads by serving
// `103` responses with `Link` headers from the final response. Refer to
// [Early Hints](https://developers.cloudflare.com/cache/about/early-hints) for
// more information.
type ZoneSettingEditParamsItemsEarlyHints struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsEarlyHintsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsEarlyHintsValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsEarlyHints) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsEarlyHints) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsEarlyHintsID string

const (
	ZoneSettingEditParamsItemsEarlyHintsIDEarlyHints ZoneSettingEditParamsItemsEarlyHintsID = "early_hints"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsEarlyHintsEditable bool

const (
	ZoneSettingEditParamsItemsEarlyHintsEditableTrue  ZoneSettingEditParamsItemsEarlyHintsEditable = true
	ZoneSettingEditParamsItemsEarlyHintsEditableFalse ZoneSettingEditParamsItemsEarlyHintsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsEarlyHintsValue string

const (
	ZoneSettingEditParamsItemsEarlyHintsValueOn  ZoneSettingEditParamsItemsEarlyHintsValue = "on"
	ZoneSettingEditParamsItemsEarlyHintsValueOff ZoneSettingEditParamsItemsEarlyHintsValue = "off"
)

// Time (in seconds) that a resource will be ensured to remain on Cloudflare's
// cache servers.
type ZoneSettingEditParamsItemsEdgeCacheTtl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsEdgeCacheTtlID] `json:"id"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value param.Field[ZoneSettingEditParamsItemsEdgeCacheTtlValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsEdgeCacheTtl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsEdgeCacheTtl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsEdgeCacheTtlID string

const (
	ZoneSettingEditParamsItemsEdgeCacheTtlIDEdgeCacheTtl ZoneSettingEditParamsItemsEdgeCacheTtlID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsEdgeCacheTtlEditable bool

const (
	ZoneSettingEditParamsItemsEdgeCacheTtlEditableTrue  ZoneSettingEditParamsItemsEdgeCacheTtlEditable = true
	ZoneSettingEditParamsItemsEdgeCacheTtlEditableFalse ZoneSettingEditParamsItemsEdgeCacheTtlEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingEditParamsItemsEdgeCacheTtlValue float64

const (
	ZoneSettingEditParamsItemsEdgeCacheTtlValue30     ZoneSettingEditParamsItemsEdgeCacheTtlValue = 30
	ZoneSettingEditParamsItemsEdgeCacheTtlValue60     ZoneSettingEditParamsItemsEdgeCacheTtlValue = 60
	ZoneSettingEditParamsItemsEdgeCacheTtlValue300    ZoneSettingEditParamsItemsEdgeCacheTtlValue = 300
	ZoneSettingEditParamsItemsEdgeCacheTtlValue1200   ZoneSettingEditParamsItemsEdgeCacheTtlValue = 1200
	ZoneSettingEditParamsItemsEdgeCacheTtlValue1800   ZoneSettingEditParamsItemsEdgeCacheTtlValue = 1800
	ZoneSettingEditParamsItemsEdgeCacheTtlValue3600   ZoneSettingEditParamsItemsEdgeCacheTtlValue = 3600
	ZoneSettingEditParamsItemsEdgeCacheTtlValue7200   ZoneSettingEditParamsItemsEdgeCacheTtlValue = 7200
	ZoneSettingEditParamsItemsEdgeCacheTtlValue10800  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 10800
	ZoneSettingEditParamsItemsEdgeCacheTtlValue14400  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 14400
	ZoneSettingEditParamsItemsEdgeCacheTtlValue18000  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 18000
	ZoneSettingEditParamsItemsEdgeCacheTtlValue28800  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 28800
	ZoneSettingEditParamsItemsEdgeCacheTtlValue43200  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 43200
	ZoneSettingEditParamsItemsEdgeCacheTtlValue57600  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 57600
	ZoneSettingEditParamsItemsEdgeCacheTtlValue72000  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 72000
	ZoneSettingEditParamsItemsEdgeCacheTtlValue86400  ZoneSettingEditParamsItemsEdgeCacheTtlValue = 86400
	ZoneSettingEditParamsItemsEdgeCacheTtlValue172800 ZoneSettingEditParamsItemsEdgeCacheTtlValue = 172800
	ZoneSettingEditParamsItemsEdgeCacheTtlValue259200 ZoneSettingEditParamsItemsEdgeCacheTtlValue = 259200
	ZoneSettingEditParamsItemsEdgeCacheTtlValue345600 ZoneSettingEditParamsItemsEdgeCacheTtlValue = 345600
	ZoneSettingEditParamsItemsEdgeCacheTtlValue432000 ZoneSettingEditParamsItemsEdgeCacheTtlValue = 432000
	ZoneSettingEditParamsItemsEdgeCacheTtlValue518400 ZoneSettingEditParamsItemsEdgeCacheTtlValue = 518400
	ZoneSettingEditParamsItemsEdgeCacheTtlValue604800 ZoneSettingEditParamsItemsEdgeCacheTtlValue = 604800
)

// Encrypt email adresses on your web page from bots, while keeping them visible to
// humans. (https://support.cloudflare.com/hc/en-us/articles/200170016).
type ZoneSettingEditParamsItemsEmailObfuscation struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsEmailObfuscationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsEmailObfuscationValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsEmailObfuscation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsEmailObfuscation) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsEmailObfuscationID string

const (
	ZoneSettingEditParamsItemsEmailObfuscationIDEmailObfuscation ZoneSettingEditParamsItemsEmailObfuscationID = "email_obfuscation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsEmailObfuscationEditable bool

const (
	ZoneSettingEditParamsItemsEmailObfuscationEditableTrue  ZoneSettingEditParamsItemsEmailObfuscationEditable = true
	ZoneSettingEditParamsItemsEmailObfuscationEditableFalse ZoneSettingEditParamsItemsEmailObfuscationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsEmailObfuscationValue string

const (
	ZoneSettingEditParamsItemsEmailObfuscationValueOn  ZoneSettingEditParamsItemsEmailObfuscationValue = "on"
	ZoneSettingEditParamsItemsEmailObfuscationValueOff ZoneSettingEditParamsItemsEmailObfuscationValue = "off"
)

// HTTP/2 Edge Prioritization optimises the delivery of resources served through
// HTTP/2 to improve page load performance. It also supports fine control of
// content delivery when used in conjunction with Workers.
type ZoneSettingEditParamsItemsH2Prioritization struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsH2PrioritizationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsH2PrioritizationValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsH2Prioritization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsH2Prioritization) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsH2PrioritizationID string

const (
	ZoneSettingEditParamsItemsH2PrioritizationIDH2Prioritization ZoneSettingEditParamsItemsH2PrioritizationID = "h2_prioritization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsH2PrioritizationEditable bool

const (
	ZoneSettingEditParamsItemsH2PrioritizationEditableTrue  ZoneSettingEditParamsItemsH2PrioritizationEditable = true
	ZoneSettingEditParamsItemsH2PrioritizationEditableFalse ZoneSettingEditParamsItemsH2PrioritizationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsH2PrioritizationValue string

const (
	ZoneSettingEditParamsItemsH2PrioritizationValueOn     ZoneSettingEditParamsItemsH2PrioritizationValue = "on"
	ZoneSettingEditParamsItemsH2PrioritizationValueOff    ZoneSettingEditParamsItemsH2PrioritizationValue = "off"
	ZoneSettingEditParamsItemsH2PrioritizationValueCustom ZoneSettingEditParamsItemsH2PrioritizationValue = "custom"
)

// When enabled, the Hotlink Protection option ensures that other sites cannot suck
// up your bandwidth by building pages that use images hosted on your site. Anytime
// a request for an image on your site hits Cloudflare, we check to ensure that
// it's not another site requesting them. People will still be able to download and
// view images from your page, but other sites won't be able to steal them for use
// on their own pages.
// (https://support.cloudflare.com/hc/en-us/articles/200170026).
type ZoneSettingEditParamsItemsHotlinkProtection struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsHotlinkProtectionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsHotlinkProtectionValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsHotlinkProtection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsHotlinkProtection) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsHotlinkProtectionID string

const (
	ZoneSettingEditParamsItemsHotlinkProtectionIDHotlinkProtection ZoneSettingEditParamsItemsHotlinkProtectionID = "hotlink_protection"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsHotlinkProtectionEditable bool

const (
	ZoneSettingEditParamsItemsHotlinkProtectionEditableTrue  ZoneSettingEditParamsItemsHotlinkProtectionEditable = true
	ZoneSettingEditParamsItemsHotlinkProtectionEditableFalse ZoneSettingEditParamsItemsHotlinkProtectionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsHotlinkProtectionValue string

const (
	ZoneSettingEditParamsItemsHotlinkProtectionValueOn  ZoneSettingEditParamsItemsHotlinkProtectionValue = "on"
	ZoneSettingEditParamsItemsHotlinkProtectionValueOff ZoneSettingEditParamsItemsHotlinkProtectionValue = "off"
)

// HTTP2 enabled for this zone.
type ZoneSettingEditParamsItemsHttp2 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsHttp2ID] `json:"id"`
	// Value of the HTTP2 setting.
	Value param.Field[ZoneSettingEditParamsItemsHttp2Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsHttp2) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsHttp2) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsHttp2ID string

const (
	ZoneSettingEditParamsItemsHttp2IDHttp2 ZoneSettingEditParamsItemsHttp2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsHttp2Editable bool

const (
	ZoneSettingEditParamsItemsHttp2EditableTrue  ZoneSettingEditParamsItemsHttp2Editable = true
	ZoneSettingEditParamsItemsHttp2EditableFalse ZoneSettingEditParamsItemsHttp2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingEditParamsItemsHttp2Value string

const (
	ZoneSettingEditParamsItemsHttp2ValueOn  ZoneSettingEditParamsItemsHttp2Value = "on"
	ZoneSettingEditParamsItemsHttp2ValueOff ZoneSettingEditParamsItemsHttp2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingEditParamsItemsHttp3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsHttp3ID] `json:"id"`
	// Value of the HTTP3 setting.
	Value param.Field[ZoneSettingEditParamsItemsHttp3Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsHttp3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsHttp3) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsHttp3ID string

const (
	ZoneSettingEditParamsItemsHttp3IDHttp3 ZoneSettingEditParamsItemsHttp3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsHttp3Editable bool

const (
	ZoneSettingEditParamsItemsHttp3EditableTrue  ZoneSettingEditParamsItemsHttp3Editable = true
	ZoneSettingEditParamsItemsHttp3EditableFalse ZoneSettingEditParamsItemsHttp3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingEditParamsItemsHttp3Value string

const (
	ZoneSettingEditParamsItemsHttp3ValueOn  ZoneSettingEditParamsItemsHttp3Value = "on"
	ZoneSettingEditParamsItemsHttp3ValueOff ZoneSettingEditParamsItemsHttp3Value = "off"
)

// Image Resizing provides on-demand resizing, conversion and optimisation for
// images served through Cloudflare's network. Refer to the
// [Image Resizing documentation](https://developers.cloudflare.com/images/) for
// more information.
type ZoneSettingEditParamsItemsImageResizing struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsImageResizingID] `json:"id"`
	// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
	Value param.Field[ZoneSettingEditParamsItemsImageResizingValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsImageResizing) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsImageResizing) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsImageResizingID string

const (
	ZoneSettingEditParamsItemsImageResizingIDImageResizing ZoneSettingEditParamsItemsImageResizingID = "image_resizing"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsImageResizingEditable bool

const (
	ZoneSettingEditParamsItemsImageResizingEditableTrue  ZoneSettingEditParamsItemsImageResizingEditable = true
	ZoneSettingEditParamsItemsImageResizingEditableFalse ZoneSettingEditParamsItemsImageResizingEditable = false
)

// Whether the feature is enabled, disabled, or enabled in `open proxy` mode.
type ZoneSettingEditParamsItemsImageResizingValue string

const (
	ZoneSettingEditParamsItemsImageResizingValueOn   ZoneSettingEditParamsItemsImageResizingValue = "on"
	ZoneSettingEditParamsItemsImageResizingValueOff  ZoneSettingEditParamsItemsImageResizingValue = "off"
	ZoneSettingEditParamsItemsImageResizingValueOpen ZoneSettingEditParamsItemsImageResizingValue = "open"
)

// Enable IP Geolocation to have Cloudflare geolocate visitors to your website and
// pass the country code to you.
// (https://support.cloudflare.com/hc/en-us/articles/200168236).
type ZoneSettingEditParamsItemsIPGeolocation struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsIPGeolocationID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsIPGeolocationValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsIPGeolocation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsIPGeolocation) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsIPGeolocationID string

const (
	ZoneSettingEditParamsItemsIPGeolocationIDIPGeolocation ZoneSettingEditParamsItemsIPGeolocationID = "ip_geolocation"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsIPGeolocationEditable bool

const (
	ZoneSettingEditParamsItemsIPGeolocationEditableTrue  ZoneSettingEditParamsItemsIPGeolocationEditable = true
	ZoneSettingEditParamsItemsIPGeolocationEditableFalse ZoneSettingEditParamsItemsIPGeolocationEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsIPGeolocationValue string

const (
	ZoneSettingEditParamsItemsIPGeolocationValueOn  ZoneSettingEditParamsItemsIPGeolocationValue = "on"
	ZoneSettingEditParamsItemsIPGeolocationValueOff ZoneSettingEditParamsItemsIPGeolocationValue = "off"
)

// Enable IPv6 on all subdomains that are Cloudflare enabled.
// (https://support.cloudflare.com/hc/en-us/articles/200168586).
type ZoneSettingEditParamsItemsIpv6 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsIpv6ID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsIpv6Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsIpv6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsIpv6) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsIpv6ID string

const (
	ZoneSettingEditParamsItemsIpv6IDIpv6 ZoneSettingEditParamsItemsIpv6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsIpv6Editable bool

const (
	ZoneSettingEditParamsItemsIpv6EditableTrue  ZoneSettingEditParamsItemsIpv6Editable = true
	ZoneSettingEditParamsItemsIpv6EditableFalse ZoneSettingEditParamsItemsIpv6Editable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsIpv6Value string

const (
	ZoneSettingEditParamsItemsIpv6ValueOff ZoneSettingEditParamsItemsIpv6Value = "off"
	ZoneSettingEditParamsItemsIpv6ValueOn  ZoneSettingEditParamsItemsIpv6Value = "on"
)

// Maximum size of an allowable upload.
type ZoneSettingEditParamsItemsMaxUpload struct {
	// identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsMaxUploadID] `json:"id"`
	// Value of the zone setting. Notes: The size depends on the plan level of the
	// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
	Value param.Field[ZoneSettingEditParamsItemsMaxUploadValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsMaxUpload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsMaxUpload) implementsZoneSettingEditParamsItem() {}

// identifier of the zone setting.
type ZoneSettingEditParamsItemsMaxUploadID string

const (
	ZoneSettingEditParamsItemsMaxUploadIDMaxUpload ZoneSettingEditParamsItemsMaxUploadID = "max_upload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsMaxUploadEditable bool

const (
	ZoneSettingEditParamsItemsMaxUploadEditableTrue  ZoneSettingEditParamsItemsMaxUploadEditable = true
	ZoneSettingEditParamsItemsMaxUploadEditableFalse ZoneSettingEditParamsItemsMaxUploadEditable = false
)

// Value of the zone setting. Notes: The size depends on the plan level of the
// zone. (Enterprise = 500, Business = 200, Pro = 100, Free = 100)
type ZoneSettingEditParamsItemsMaxUploadValue float64

const (
	ZoneSettingEditParamsItemsMaxUploadValue100 ZoneSettingEditParamsItemsMaxUploadValue = 100
	ZoneSettingEditParamsItemsMaxUploadValue200 ZoneSettingEditParamsItemsMaxUploadValue = 200
	ZoneSettingEditParamsItemsMaxUploadValue500 ZoneSettingEditParamsItemsMaxUploadValue = 500
)

// Only accepts HTTPS requests that use at least the TLS protocol version
// specified. For example, if TLS 1.1 is selected, TLS 1.0 connections will be
// rejected, while 1.1, 1.2, and 1.3 (if enabled) will be permitted.
type ZoneSettingEditParamsItemsMinTlsVersion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsMinTlsVersionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsMinTlsVersionValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsMinTlsVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsMinTlsVersion) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsMinTlsVersionID string

const (
	ZoneSettingEditParamsItemsMinTlsVersionIDMinTlsVersion ZoneSettingEditParamsItemsMinTlsVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsMinTlsVersionEditable bool

const (
	ZoneSettingEditParamsItemsMinTlsVersionEditableTrue  ZoneSettingEditParamsItemsMinTlsVersionEditable = true
	ZoneSettingEditParamsItemsMinTlsVersionEditableFalse ZoneSettingEditParamsItemsMinTlsVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsMinTlsVersionValue string

const (
	ZoneSettingEditParamsItemsMinTlsVersionValue1_0 ZoneSettingEditParamsItemsMinTlsVersionValue = "1.0"
	ZoneSettingEditParamsItemsMinTlsVersionValue1_1 ZoneSettingEditParamsItemsMinTlsVersionValue = "1.1"
	ZoneSettingEditParamsItemsMinTlsVersionValue1_2 ZoneSettingEditParamsItemsMinTlsVersionValue = "1.2"
	ZoneSettingEditParamsItemsMinTlsVersionValue1_3 ZoneSettingEditParamsItemsMinTlsVersionValue = "1.3"
)

// Automatically minify certain assets for your website. Refer to
// [Using Cloudflare Auto Minify](https://support.cloudflare.com/hc/en-us/articles/200168196)
// for more information.
type ZoneSettingEditParamsItemsMinify struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsMinifyID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsMinifyValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsMinify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsMinify) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsMinifyID string

const (
	ZoneSettingEditParamsItemsMinifyIDMinify ZoneSettingEditParamsItemsMinifyID = "minify"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsMinifyEditable bool

const (
	ZoneSettingEditParamsItemsMinifyEditableTrue  ZoneSettingEditParamsItemsMinifyEditable = true
	ZoneSettingEditParamsItemsMinifyEditableFalse ZoneSettingEditParamsItemsMinifyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsMinifyValue struct {
	// Automatically minify all CSS files for your website.
	Css param.Field[ZoneSettingEditParamsItemsMinifyValueCss] `json:"css"`
	// Automatically minify all HTML files for your website.
	HTML param.Field[ZoneSettingEditParamsItemsMinifyValueHTML] `json:"html"`
	// Automatically minify all JavaScript files for your website.
	Js param.Field[ZoneSettingEditParamsItemsMinifyValueJs] `json:"js"`
}

func (r ZoneSettingEditParamsItemsMinifyValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Automatically minify all CSS files for your website.
type ZoneSettingEditParamsItemsMinifyValueCss string

const (
	ZoneSettingEditParamsItemsMinifyValueCssOn  ZoneSettingEditParamsItemsMinifyValueCss = "on"
	ZoneSettingEditParamsItemsMinifyValueCssOff ZoneSettingEditParamsItemsMinifyValueCss = "off"
)

// Automatically minify all HTML files for your website.
type ZoneSettingEditParamsItemsMinifyValueHTML string

const (
	ZoneSettingEditParamsItemsMinifyValueHTMLOn  ZoneSettingEditParamsItemsMinifyValueHTML = "on"
	ZoneSettingEditParamsItemsMinifyValueHTMLOff ZoneSettingEditParamsItemsMinifyValueHTML = "off"
)

// Automatically minify all JavaScript files for your website.
type ZoneSettingEditParamsItemsMinifyValueJs string

const (
	ZoneSettingEditParamsItemsMinifyValueJsOn  ZoneSettingEditParamsItemsMinifyValueJs = "on"
	ZoneSettingEditParamsItemsMinifyValueJsOff ZoneSettingEditParamsItemsMinifyValueJs = "off"
)

// Automatically optimize image loading for website visitors on mobile devices.
// Refer to
// [our blog post](http://blog.cloudflare.com/mirage2-solving-mobile-speed) for
// more information.
type ZoneSettingEditParamsItemsMirage struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsMirageID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsMirageValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsMirage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsMirage) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsMirageID string

const (
	ZoneSettingEditParamsItemsMirageIDMirage ZoneSettingEditParamsItemsMirageID = "mirage"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsMirageEditable bool

const (
	ZoneSettingEditParamsItemsMirageEditableTrue  ZoneSettingEditParamsItemsMirageEditable = true
	ZoneSettingEditParamsItemsMirageEditableFalse ZoneSettingEditParamsItemsMirageEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsMirageValue string

const (
	ZoneSettingEditParamsItemsMirageValueOn  ZoneSettingEditParamsItemsMirageValue = "on"
	ZoneSettingEditParamsItemsMirageValueOff ZoneSettingEditParamsItemsMirageValue = "off"
)

// Automatically redirect visitors on mobile devices to a mobile-optimized
// subdomain. Refer to
// [Understanding Cloudflare Mobile Redirect](https://support.cloudflare.com/hc/articles/200168336)
// for more information.
type ZoneSettingEditParamsItemsMobileRedirect struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsMobileRedirectID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsMobileRedirectValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsMobileRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsMobileRedirect) implementsZoneSettingEditParamsItem() {}

// Identifier of the zone setting.
type ZoneSettingEditParamsItemsMobileRedirectID string

const (
	ZoneSettingEditParamsItemsMobileRedirectIDMobileRedirect ZoneSettingEditParamsItemsMobileRedirectID = "mobile_redirect"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsMobileRedirectEditable bool

const (
	ZoneSettingEditParamsItemsMobileRedirectEditableTrue  ZoneSettingEditParamsItemsMobileRedirectEditable = true
	ZoneSettingEditParamsItemsMobileRedirectEditableFalse ZoneSettingEditParamsItemsMobileRedirectEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsMobileRedirectValue struct {
	// Which subdomain prefix you wish to redirect visitors on mobile devices to
	// (subdomain must already exist).
	MobileSubdomain param.Field[string] `json:"mobile_subdomain"`
	// Whether or not mobile redirect is enabled.
	Status param.Field[ZoneSettingEditParamsItemsMobileRedirectValueStatus] `json:"status"`
	// Whether to drop the current page path and redirect to the mobile subdomain URL
	// root, or keep the path and redirect to the same page on the mobile subdomain.
	StripUri param.Field[bool] `json:"strip_uri"`
}

func (r ZoneSettingEditParamsItemsMobileRedirectValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not mobile redirect is enabled.
type ZoneSettingEditParamsItemsMobileRedirectValueStatus string

const (
	ZoneSettingEditParamsItemsMobileRedirectValueStatusOn  ZoneSettingEditParamsItemsMobileRedirectValueStatus = "on"
	ZoneSettingEditParamsItemsMobileRedirectValueStatusOff ZoneSettingEditParamsItemsMobileRedirectValueStatus = "off"
)

// Enable Network Error Logging reporting on your zone. (Beta)
type ZoneSettingEditParamsItemsNel struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsNelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsNelValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsNel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsNel) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsNelID string

const (
	ZoneSettingEditParamsItemsNelIDNel ZoneSettingEditParamsItemsNelID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsNelEditable bool

const (
	ZoneSettingEditParamsItemsNelEditableTrue  ZoneSettingEditParamsItemsNelEditable = true
	ZoneSettingEditParamsItemsNelEditableFalse ZoneSettingEditParamsItemsNelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsNelValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsNelValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables the Opportunistic Encryption feature for a zone.
type ZoneSettingEditParamsItemsOpportunisticEncryption struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsOpportunisticEncryptionID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsOpportunisticEncryptionValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsOpportunisticEncryption) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsOpportunisticEncryption) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsOpportunisticEncryptionID string

const (
	ZoneSettingEditParamsItemsOpportunisticEncryptionIDOpportunisticEncryption ZoneSettingEditParamsItemsOpportunisticEncryptionID = "opportunistic_encryption"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsOpportunisticEncryptionEditable bool

const (
	ZoneSettingEditParamsItemsOpportunisticEncryptionEditableTrue  ZoneSettingEditParamsItemsOpportunisticEncryptionEditable = true
	ZoneSettingEditParamsItemsOpportunisticEncryptionEditableFalse ZoneSettingEditParamsItemsOpportunisticEncryptionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsOpportunisticEncryptionValue string

const (
	ZoneSettingEditParamsItemsOpportunisticEncryptionValueOn  ZoneSettingEditParamsItemsOpportunisticEncryptionValue = "on"
	ZoneSettingEditParamsItemsOpportunisticEncryptionValueOff ZoneSettingEditParamsItemsOpportunisticEncryptionValue = "off"
)

// Add an Alt-Svc header to all legitimate requests from Tor, allowing the
// connection to use our onion services instead of exit nodes.
type ZoneSettingEditParamsItemsOpportunisticOnion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsOpportunisticOnionID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsOpportunisticOnionValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsOpportunisticOnion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsOpportunisticOnion) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsOpportunisticOnionID string

const (
	ZoneSettingEditParamsItemsOpportunisticOnionIDOpportunisticOnion ZoneSettingEditParamsItemsOpportunisticOnionID = "opportunistic_onion"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsOpportunisticOnionEditable bool

const (
	ZoneSettingEditParamsItemsOpportunisticOnionEditableTrue  ZoneSettingEditParamsItemsOpportunisticOnionEditable = true
	ZoneSettingEditParamsItemsOpportunisticOnionEditableFalse ZoneSettingEditParamsItemsOpportunisticOnionEditable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsOpportunisticOnionValue string

const (
	ZoneSettingEditParamsItemsOpportunisticOnionValueOn  ZoneSettingEditParamsItemsOpportunisticOnionValue = "on"
	ZoneSettingEditParamsItemsOpportunisticOnionValueOff ZoneSettingEditParamsItemsOpportunisticOnionValue = "off"
)

// Orange to Orange (O2O) allows zones on Cloudflare to CNAME to other zones also
// on Cloudflare.
type ZoneSettingEditParamsItemsOrangeToOrange struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsOrangeToOrangeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsOrangeToOrangeValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsOrangeToOrange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsOrangeToOrange) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsOrangeToOrangeID string

const (
	ZoneSettingEditParamsItemsOrangeToOrangeIDOrangeToOrange ZoneSettingEditParamsItemsOrangeToOrangeID = "orange_to_orange"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsOrangeToOrangeEditable bool

const (
	ZoneSettingEditParamsItemsOrangeToOrangeEditableTrue  ZoneSettingEditParamsItemsOrangeToOrangeEditable = true
	ZoneSettingEditParamsItemsOrangeToOrangeEditableFalse ZoneSettingEditParamsItemsOrangeToOrangeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsOrangeToOrangeValue string

const (
	ZoneSettingEditParamsItemsOrangeToOrangeValueOn  ZoneSettingEditParamsItemsOrangeToOrangeValue = "on"
	ZoneSettingEditParamsItemsOrangeToOrangeValueOff ZoneSettingEditParamsItemsOrangeToOrangeValue = "off"
)

// Cloudflare will proxy customer error pages on any 502,504 errors on origin
// server instead of showing a default Cloudflare error page. This does not apply
// to 522 errors and is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsOriginErrorPagePassThru struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsOriginErrorPagePassThruID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsOriginErrorPagePassThruValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsOriginErrorPagePassThru) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsOriginErrorPagePassThru) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsOriginErrorPagePassThruID string

const (
	ZoneSettingEditParamsItemsOriginErrorPagePassThruIDOriginErrorPagePassThru ZoneSettingEditParamsItemsOriginErrorPagePassThruID = "origin_error_page_pass_thru"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsOriginErrorPagePassThruEditable bool

const (
	ZoneSettingEditParamsItemsOriginErrorPagePassThruEditableTrue  ZoneSettingEditParamsItemsOriginErrorPagePassThruEditable = true
	ZoneSettingEditParamsItemsOriginErrorPagePassThruEditableFalse ZoneSettingEditParamsItemsOriginErrorPagePassThruEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsOriginErrorPagePassThruValue string

const (
	ZoneSettingEditParamsItemsOriginErrorPagePassThruValueOn  ZoneSettingEditParamsItemsOriginErrorPagePassThruValue = "on"
	ZoneSettingEditParamsItemsOriginErrorPagePassThruValueOff ZoneSettingEditParamsItemsOriginErrorPagePassThruValue = "off"
)

type ZoneSettingEditParamsItemsOriginMaxHTTPVersion struct {
	// Identifier of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsOriginMaxHTTPVersionID] `json:"id,required"`
}

func (r ZoneSettingEditParamsItemsOriginMaxHTTPVersion) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsOriginMaxHTTPVersion) implementsZoneSettingEditParamsItem() {}

// Identifier of the zone setting.
type ZoneSettingEditParamsItemsOriginMaxHTTPVersionID string

const (
	ZoneSettingEditParamsItemsOriginMaxHTTPVersionIDOriginMaxHTTPVersion ZoneSettingEditParamsItemsOriginMaxHTTPVersionID = "origin_max_http_version"
)

// Removes metadata and compresses your images for faster page load times. Basic
// (Lossless): Reduce the size of PNG, JPEG, and GIF files - no impact on visual
// quality. Basic + JPEG (Lossy): Further reduce the size of JPEG files for faster
// image loading. Larger JPEGs are converted to progressive images, loading a
// lower-resolution image first and ending in a higher-resolution version. Not
// recommended for hi-res photography sites.
type ZoneSettingEditParamsItemsPolish struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsPolishID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsPolishValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsPolish) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsPolish) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsPolishID string

const (
	ZoneSettingEditParamsItemsPolishIDPolish ZoneSettingEditParamsItemsPolishID = "polish"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsPolishEditable bool

const (
	ZoneSettingEditParamsItemsPolishEditableTrue  ZoneSettingEditParamsItemsPolishEditable = true
	ZoneSettingEditParamsItemsPolishEditableFalse ZoneSettingEditParamsItemsPolishEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsPolishValue string

const (
	ZoneSettingEditParamsItemsPolishValueOff      ZoneSettingEditParamsItemsPolishValue = "off"
	ZoneSettingEditParamsItemsPolishValueLossless ZoneSettingEditParamsItemsPolishValue = "lossless"
	ZoneSettingEditParamsItemsPolishValueLossy    ZoneSettingEditParamsItemsPolishValue = "lossy"
)

// Cloudflare will prefetch any URLs that are included in the response headers.
// This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsPrefetchPreload struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsPrefetchPreloadID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsPrefetchPreloadValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsPrefetchPreload) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsPrefetchPreload) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsPrefetchPreloadID string

const (
	ZoneSettingEditParamsItemsPrefetchPreloadIDPrefetchPreload ZoneSettingEditParamsItemsPrefetchPreloadID = "prefetch_preload"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsPrefetchPreloadEditable bool

const (
	ZoneSettingEditParamsItemsPrefetchPreloadEditableTrue  ZoneSettingEditParamsItemsPrefetchPreloadEditable = true
	ZoneSettingEditParamsItemsPrefetchPreloadEditableFalse ZoneSettingEditParamsItemsPrefetchPreloadEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsPrefetchPreloadValue string

const (
	ZoneSettingEditParamsItemsPrefetchPreloadValueOn  ZoneSettingEditParamsItemsPrefetchPreloadValue = "on"
	ZoneSettingEditParamsItemsPrefetchPreloadValueOff ZoneSettingEditParamsItemsPrefetchPreloadValue = "off"
)

// Privacy Pass is a browser extension developed by the Privacy Pass Team to
// improve the browsing experience for your visitors. Enabling Privacy Pass will
// reduce the number of CAPTCHAs shown to your visitors.
// (https://support.cloudflare.com/hc/en-us/articles/115001992652-Privacy-Pass).
type ZoneSettingEditParamsItemsPrivacyPass struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsPrivacyPassID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsPrivacyPassValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsPrivacyPass) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsPrivacyPass) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsPrivacyPassID string

const (
	ZoneSettingEditParamsItemsPrivacyPassIDPrivacyPass ZoneSettingEditParamsItemsPrivacyPassID = "privacy_pass"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsPrivacyPassEditable bool

const (
	ZoneSettingEditParamsItemsPrivacyPassEditableTrue  ZoneSettingEditParamsItemsPrivacyPassEditable = true
	ZoneSettingEditParamsItemsPrivacyPassEditableFalse ZoneSettingEditParamsItemsPrivacyPassEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsPrivacyPassValue string

const (
	ZoneSettingEditParamsItemsPrivacyPassValueOn  ZoneSettingEditParamsItemsPrivacyPassValue = "on"
	ZoneSettingEditParamsItemsPrivacyPassValueOff ZoneSettingEditParamsItemsPrivacyPassValue = "off"
)

// Maximum time between two read operations from origin.
type ZoneSettingEditParamsItemsProxyReadTimeout struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsProxyReadTimeoutID] `json:"id"`
	// Value of the zone setting. Notes: Value must be between 1 and 6000
	Value param.Field[float64] `json:"value"`
}

func (r ZoneSettingEditParamsItemsProxyReadTimeout) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsProxyReadTimeout) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsProxyReadTimeoutID string

const (
	ZoneSettingEditParamsItemsProxyReadTimeoutIDProxyReadTimeout ZoneSettingEditParamsItemsProxyReadTimeoutID = "proxy_read_timeout"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsProxyReadTimeoutEditable bool

const (
	ZoneSettingEditParamsItemsProxyReadTimeoutEditableTrue  ZoneSettingEditParamsItemsProxyReadTimeoutEditable = true
	ZoneSettingEditParamsItemsProxyReadTimeoutEditableFalse ZoneSettingEditParamsItemsProxyReadTimeoutEditable = false
)

// The value set for the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsPseudoIpv4 struct {
	// Value of the Pseudo IPv4 setting.
	ID param.Field[ZoneSettingEditParamsItemsPseudoIpv4ID] `json:"id"`
	// Value of the Pseudo IPv4 setting.
	Value param.Field[ZoneSettingEditParamsItemsPseudoIpv4Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsPseudoIpv4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsPseudoIpv4) implementsZoneSettingEditParamsItem() {}

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsPseudoIpv4ID string

const (
	ZoneSettingEditParamsItemsPseudoIpv4IDPseudoIpv4 ZoneSettingEditParamsItemsPseudoIpv4ID = "pseudo_ipv4"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsPseudoIpv4Editable bool

const (
	ZoneSettingEditParamsItemsPseudoIpv4EditableTrue  ZoneSettingEditParamsItemsPseudoIpv4Editable = true
	ZoneSettingEditParamsItemsPseudoIpv4EditableFalse ZoneSettingEditParamsItemsPseudoIpv4Editable = false
)

// Value of the Pseudo IPv4 setting.
type ZoneSettingEditParamsItemsPseudoIpv4Value string

const (
	ZoneSettingEditParamsItemsPseudoIpv4ValueOff             ZoneSettingEditParamsItemsPseudoIpv4Value = "off"
	ZoneSettingEditParamsItemsPseudoIpv4ValueAddHeader       ZoneSettingEditParamsItemsPseudoIpv4Value = "add_header"
	ZoneSettingEditParamsItemsPseudoIpv4ValueOverwriteHeader ZoneSettingEditParamsItemsPseudoIpv4Value = "overwrite_header"
)

// Enables or disables buffering of responses from the proxied server. Cloudflare
// may buffer the whole payload to deliver it at once to the client versus allowing
// it to be delivered in chunks. By default, the proxied server streams directly
// and is not buffered by Cloudflare. This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsResponseBuffering struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsResponseBufferingID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsResponseBufferingValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsResponseBuffering) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsResponseBuffering) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsResponseBufferingID string

const (
	ZoneSettingEditParamsItemsResponseBufferingIDResponseBuffering ZoneSettingEditParamsItemsResponseBufferingID = "response_buffering"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsResponseBufferingEditable bool

const (
	ZoneSettingEditParamsItemsResponseBufferingEditableTrue  ZoneSettingEditParamsItemsResponseBufferingEditable = true
	ZoneSettingEditParamsItemsResponseBufferingEditableFalse ZoneSettingEditParamsItemsResponseBufferingEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsResponseBufferingValue string

const (
	ZoneSettingEditParamsItemsResponseBufferingValueOn  ZoneSettingEditParamsItemsResponseBufferingValue = "on"
	ZoneSettingEditParamsItemsResponseBufferingValueOff ZoneSettingEditParamsItemsResponseBufferingValue = "off"
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
type ZoneSettingEditParamsItemsRocketLoader struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsRocketLoaderID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsRocketLoaderValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsRocketLoader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsRocketLoader) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsRocketLoaderID string

const (
	ZoneSettingEditParamsItemsRocketLoaderIDRocketLoader ZoneSettingEditParamsItemsRocketLoaderID = "rocket_loader"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsRocketLoaderEditable bool

const (
	ZoneSettingEditParamsItemsRocketLoaderEditableTrue  ZoneSettingEditParamsItemsRocketLoaderEditable = true
	ZoneSettingEditParamsItemsRocketLoaderEditableFalse ZoneSettingEditParamsItemsRocketLoaderEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsRocketLoaderValue string

const (
	ZoneSettingEditParamsItemsRocketLoaderValueOn  ZoneSettingEditParamsItemsRocketLoaderValue = "on"
	ZoneSettingEditParamsItemsRocketLoaderValueOff ZoneSettingEditParamsItemsRocketLoaderValue = "off"
)

// [Automatic Platform Optimization for WordPress](https://developers.cloudflare.com/automatic-platform-optimization/)
// serves your WordPress site from Cloudflare's edge network and caches third-party
// fonts.
type ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimization struct {
	// ID of the zone setting.
	ID    param.Field[ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationID]    `json:"id"`
	Value param.Field[ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimization) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimization) implementsZoneSettingEditParamsItem() {
}

// ID of the zone setting.
type ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationID string

const (
	ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationIDAutomaticPlatformOptimization ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationID = "automatic_platform_optimization"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationEditable bool

const (
	ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationEditableTrue  ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationEditable = true
	ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationEditableFalse ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationEditable = false
)

type ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationValue struct {
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

func (r ZoneSettingEditParamsItemsSchemasAutomaticPlatformOptimizationValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloudflare security header for a zone.
type ZoneSettingEditParamsItemsSecurityHeader struct {
	// ID of the zone's security header.
	ID    param.Field[ZoneSettingEditParamsItemsSecurityHeaderID]    `json:"id"`
	Value param.Field[ZoneSettingEditParamsItemsSecurityHeaderValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsSecurityHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSecurityHeader) implementsZoneSettingEditParamsItem() {}

// ID of the zone's security header.
type ZoneSettingEditParamsItemsSecurityHeaderID string

const (
	ZoneSettingEditParamsItemsSecurityHeaderIDSecurityHeader ZoneSettingEditParamsItemsSecurityHeaderID = "security_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSecurityHeaderEditable bool

const (
	ZoneSettingEditParamsItemsSecurityHeaderEditableTrue  ZoneSettingEditParamsItemsSecurityHeaderEditable = true
	ZoneSettingEditParamsItemsSecurityHeaderEditableFalse ZoneSettingEditParamsItemsSecurityHeaderEditable = false
)

type ZoneSettingEditParamsItemsSecurityHeaderValue struct {
	// Strict Transport Security.
	StrictTransportSecurity param.Field[ZoneSettingEditParamsItemsSecurityHeaderValueStrictTransportSecurity] `json:"strict_transport_security"`
}

func (r ZoneSettingEditParamsItemsSecurityHeaderValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Strict Transport Security.
type ZoneSettingEditParamsItemsSecurityHeaderValueStrictTransportSecurity struct {
	// Whether or not strict transport security is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Include all subdomains for strict transport security.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
	// Max age in seconds of the strict transport security.
	MaxAge param.Field[float64] `json:"max_age"`
	// Whether or not to include 'X-Content-Type-Options: nosniff' header.
	Nosniff param.Field[bool] `json:"nosniff"`
}

func (r ZoneSettingEditParamsItemsSecurityHeaderValueStrictTransportSecurity) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose the appropriate security profile for your website, which will
// automatically adjust each of the security settings. If you choose to customize
// an individual security setting, the profile will become Custom.
// (https://support.cloudflare.com/hc/en-us/articles/200170056).
type ZoneSettingEditParamsItemsSecurityLevel struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsSecurityLevelID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsSecurityLevelValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsSecurityLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSecurityLevel) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsSecurityLevelID string

const (
	ZoneSettingEditParamsItemsSecurityLevelIDSecurityLevel ZoneSettingEditParamsItemsSecurityLevelID = "security_level"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSecurityLevelEditable bool

const (
	ZoneSettingEditParamsItemsSecurityLevelEditableTrue  ZoneSettingEditParamsItemsSecurityLevelEditable = true
	ZoneSettingEditParamsItemsSecurityLevelEditableFalse ZoneSettingEditParamsItemsSecurityLevelEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsSecurityLevelValue string

const (
	ZoneSettingEditParamsItemsSecurityLevelValueOff            ZoneSettingEditParamsItemsSecurityLevelValue = "off"
	ZoneSettingEditParamsItemsSecurityLevelValueEssentiallyOff ZoneSettingEditParamsItemsSecurityLevelValue = "essentially_off"
	ZoneSettingEditParamsItemsSecurityLevelValueLow            ZoneSettingEditParamsItemsSecurityLevelValue = "low"
	ZoneSettingEditParamsItemsSecurityLevelValueMedium         ZoneSettingEditParamsItemsSecurityLevelValue = "medium"
	ZoneSettingEditParamsItemsSecurityLevelValueHigh           ZoneSettingEditParamsItemsSecurityLevelValue = "high"
	ZoneSettingEditParamsItemsSecurityLevelValueUnderAttack    ZoneSettingEditParamsItemsSecurityLevelValue = "under_attack"
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
type ZoneSettingEditParamsItemsServerSideExclude struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsServerSideExcludeID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsServerSideExcludeValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsServerSideExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsServerSideExclude) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsServerSideExcludeID string

const (
	ZoneSettingEditParamsItemsServerSideExcludeIDServerSideExclude ZoneSettingEditParamsItemsServerSideExcludeID = "server_side_exclude"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsServerSideExcludeEditable bool

const (
	ZoneSettingEditParamsItemsServerSideExcludeEditableTrue  ZoneSettingEditParamsItemsServerSideExcludeEditable = true
	ZoneSettingEditParamsItemsServerSideExcludeEditableFalse ZoneSettingEditParamsItemsServerSideExcludeEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsServerSideExcludeValue string

const (
	ZoneSettingEditParamsItemsServerSideExcludeValueOn  ZoneSettingEditParamsItemsServerSideExcludeValue = "on"
	ZoneSettingEditParamsItemsServerSideExcludeValueOff ZoneSettingEditParamsItemsServerSideExcludeValue = "off"
)

// Allow SHA1 support.
type ZoneSettingEditParamsItemsSha1Support struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsSha1SupportID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsSha1SupportValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsSha1Support) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSha1Support) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsSha1SupportID string

const (
	ZoneSettingEditParamsItemsSha1SupportIDSha1Support ZoneSettingEditParamsItemsSha1SupportID = "sha1_support"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSha1SupportEditable bool

const (
	ZoneSettingEditParamsItemsSha1SupportEditableTrue  ZoneSettingEditParamsItemsSha1SupportEditable = true
	ZoneSettingEditParamsItemsSha1SupportEditableFalse ZoneSettingEditParamsItemsSha1SupportEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsSha1SupportValue string

const (
	ZoneSettingEditParamsItemsSha1SupportValueOff ZoneSettingEditParamsItemsSha1SupportValue = "off"
	ZoneSettingEditParamsItemsSha1SupportValueOn  ZoneSettingEditParamsItemsSha1SupportValue = "on"
)

// Cloudflare will treat files with the same query strings as the same file in
// cache, regardless of the order of the query strings. This is limited to
// Enterprise Zones.
type ZoneSettingEditParamsItemsSortQueryStringForCache struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsSortQueryStringForCacheID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsSortQueryStringForCacheValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsSortQueryStringForCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSortQueryStringForCache) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsSortQueryStringForCacheID string

const (
	ZoneSettingEditParamsItemsSortQueryStringForCacheIDSortQueryStringForCache ZoneSettingEditParamsItemsSortQueryStringForCacheID = "sort_query_string_for_cache"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSortQueryStringForCacheEditable bool

const (
	ZoneSettingEditParamsItemsSortQueryStringForCacheEditableTrue  ZoneSettingEditParamsItemsSortQueryStringForCacheEditable = true
	ZoneSettingEditParamsItemsSortQueryStringForCacheEditableFalse ZoneSettingEditParamsItemsSortQueryStringForCacheEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsSortQueryStringForCacheValue string

const (
	ZoneSettingEditParamsItemsSortQueryStringForCacheValueOn  ZoneSettingEditParamsItemsSortQueryStringForCacheValue = "on"
	ZoneSettingEditParamsItemsSortQueryStringForCacheValueOff ZoneSettingEditParamsItemsSortQueryStringForCacheValue = "off"
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
type ZoneSettingEditParamsItemsSsl struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsSslID] `json:"id"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZoneSettingEditParamsItemsSslValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsSsl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSsl) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsSslID string

const (
	ZoneSettingEditParamsItemsSslIDSsl ZoneSettingEditParamsItemsSslID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSslEditable bool

const (
	ZoneSettingEditParamsItemsSslEditableTrue  ZoneSettingEditParamsItemsSslEditable = true
	ZoneSettingEditParamsItemsSslEditableFalse ZoneSettingEditParamsItemsSslEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingEditParamsItemsSslValue string

const (
	ZoneSettingEditParamsItemsSslValueOff      ZoneSettingEditParamsItemsSslValue = "off"
	ZoneSettingEditParamsItemsSslValueFlexible ZoneSettingEditParamsItemsSslValue = "flexible"
	ZoneSettingEditParamsItemsSslValueFull     ZoneSettingEditParamsItemsSslValue = "full"
	ZoneSettingEditParamsItemsSslValueStrict   ZoneSettingEditParamsItemsSslValue = "strict"
)

// Enrollment in the SSL/TLS Recommender service which tries to detect and
// recommend (by sending periodic emails) the most secure SSL/TLS setting your
// origin servers support.
type ZoneSettingEditParamsItemsSslRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID param.Field[ZoneSettingEditParamsItemsSslRecommenderID] `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsSslRecommender) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsSslRecommender) implementsZoneSettingEditParamsItem() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingEditParamsItemsSslRecommenderID string

const (
	ZoneSettingEditParamsItemsSslRecommenderIDSslRecommender ZoneSettingEditParamsItemsSslRecommenderID = "ssl_recommender"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsSslRecommenderEditable bool

const (
	ZoneSettingEditParamsItemsSslRecommenderEditableTrue  ZoneSettingEditParamsItemsSslRecommenderEditable = true
	ZoneSettingEditParamsItemsSslRecommenderEditableFalse ZoneSettingEditParamsItemsSslRecommenderEditable = false
)

// Only allows TLS1.2.
type ZoneSettingEditParamsItemsTls1_2Only struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsTls1_2OnlyID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsTls1_2OnlyValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsTls1_2Only) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsTls1_2Only) implementsZoneSettingEditParamsItem() {}

// Zone setting identifier.
type ZoneSettingEditParamsItemsTls1_2OnlyID string

const (
	ZoneSettingEditParamsItemsTls1_2OnlyIDTls1_2Only ZoneSettingEditParamsItemsTls1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsTls1_2OnlyEditable bool

const (
	ZoneSettingEditParamsItemsTls1_2OnlyEditableTrue  ZoneSettingEditParamsItemsTls1_2OnlyEditable = true
	ZoneSettingEditParamsItemsTls1_2OnlyEditableFalse ZoneSettingEditParamsItemsTls1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsTls1_2OnlyValue string

const (
	ZoneSettingEditParamsItemsTls1_2OnlyValueOff ZoneSettingEditParamsItemsTls1_2OnlyValue = "off"
	ZoneSettingEditParamsItemsTls1_2OnlyValueOn  ZoneSettingEditParamsItemsTls1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditParamsItemsTls1_3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsTls1_3ID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsTls1_3Value] `json:"value"`
}

func (r ZoneSettingEditParamsItemsTls1_3) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsTls1_3) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsTls1_3ID string

const (
	ZoneSettingEditParamsItemsTls1_3IDTls1_3 ZoneSettingEditParamsItemsTls1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsTls1_3Editable bool

const (
	ZoneSettingEditParamsItemsTls1_3EditableTrue  ZoneSettingEditParamsItemsTls1_3Editable = true
	ZoneSettingEditParamsItemsTls1_3EditableFalse ZoneSettingEditParamsItemsTls1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsTls1_3Value string

const (
	ZoneSettingEditParamsItemsTls1_3ValueOn  ZoneSettingEditParamsItemsTls1_3Value = "on"
	ZoneSettingEditParamsItemsTls1_3ValueOff ZoneSettingEditParamsItemsTls1_3Value = "off"
	ZoneSettingEditParamsItemsTls1_3ValueZrt ZoneSettingEditParamsItemsTls1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditParamsItemsTlsClientAuth struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsTlsClientAuthID] `json:"id"`
	// value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsTlsClientAuthValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsTlsClientAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsTlsClientAuth) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsTlsClientAuthID string

const (
	ZoneSettingEditParamsItemsTlsClientAuthIDTlsClientAuth ZoneSettingEditParamsItemsTlsClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsTlsClientAuthEditable bool

const (
	ZoneSettingEditParamsItemsTlsClientAuthEditableTrue  ZoneSettingEditParamsItemsTlsClientAuthEditable = true
	ZoneSettingEditParamsItemsTlsClientAuthEditableFalse ZoneSettingEditParamsItemsTlsClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingEditParamsItemsTlsClientAuthValue string

const (
	ZoneSettingEditParamsItemsTlsClientAuthValueOn  ZoneSettingEditParamsItemsTlsClientAuthValue = "on"
	ZoneSettingEditParamsItemsTlsClientAuthValueOff ZoneSettingEditParamsItemsTlsClientAuthValue = "off"
)

// Allows customer to continue to use True Client IP (Akamai feature) in the
// headers we send to the origin. This is limited to Enterprise Zones.
type ZoneSettingEditParamsItemsTrueClientIPHeader struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsTrueClientIPHeaderID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsTrueClientIPHeaderValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsTrueClientIPHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsTrueClientIPHeader) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsTrueClientIPHeaderID string

const (
	ZoneSettingEditParamsItemsTrueClientIPHeaderIDTrueClientIPHeader ZoneSettingEditParamsItemsTrueClientIPHeaderID = "true_client_ip_header"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsTrueClientIPHeaderEditable bool

const (
	ZoneSettingEditParamsItemsTrueClientIPHeaderEditableTrue  ZoneSettingEditParamsItemsTrueClientIPHeaderEditable = true
	ZoneSettingEditParamsItemsTrueClientIPHeaderEditableFalse ZoneSettingEditParamsItemsTrueClientIPHeaderEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsTrueClientIPHeaderValue string

const (
	ZoneSettingEditParamsItemsTrueClientIPHeaderValueOn  ZoneSettingEditParamsItemsTrueClientIPHeaderValue = "on"
	ZoneSettingEditParamsItemsTrueClientIPHeaderValueOff ZoneSettingEditParamsItemsTrueClientIPHeaderValue = "off"
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
type ZoneSettingEditParamsItemsWaf struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsWafID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsWafValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsWaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsWaf) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsWafID string

const (
	ZoneSettingEditParamsItemsWafIDWaf ZoneSettingEditParamsItemsWafID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsWafEditable bool

const (
	ZoneSettingEditParamsItemsWafEditableTrue  ZoneSettingEditParamsItemsWafEditable = true
	ZoneSettingEditParamsItemsWafEditableFalse ZoneSettingEditParamsItemsWafEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsWafValue string

const (
	ZoneSettingEditParamsItemsWafValueOn  ZoneSettingEditParamsItemsWafValue = "on"
	ZoneSettingEditParamsItemsWafValueOff ZoneSettingEditParamsItemsWafValue = "off"
)

// When the client requesting the image supports the WebP image codec, and WebP
// offers a performance advantage over the original image format, Cloudflare will
// serve a WebP version of the original image.
type ZoneSettingEditParamsItemsWebp struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsWebpID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsWebpValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsWebp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsWebp) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsWebpID string

const (
	ZoneSettingEditParamsItemsWebpIDWebp ZoneSettingEditParamsItemsWebpID = "webp"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsWebpEditable bool

const (
	ZoneSettingEditParamsItemsWebpEditableTrue  ZoneSettingEditParamsItemsWebpEditable = true
	ZoneSettingEditParamsItemsWebpEditableFalse ZoneSettingEditParamsItemsWebpEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsWebpValue string

const (
	ZoneSettingEditParamsItemsWebpValueOff ZoneSettingEditParamsItemsWebpValue = "off"
	ZoneSettingEditParamsItemsWebpValueOn  ZoneSettingEditParamsItemsWebpValue = "on"
)

// WebSockets are open connections sustained between the client and the origin
// server. Inside a WebSockets connection, the client and the origin can pass data
// back and forth without having to reestablish sessions. This makes exchanging
// data within a WebSockets connection fast. WebSockets are often used for
// real-time applications such as live chat and gaming. For more information refer
// to
// [Can I use Cloudflare with Websockets](https://support.cloudflare.com/hc/en-us/articles/200169466-Can-I-use-Cloudflare-with-WebSockets-).
type ZoneSettingEditParamsItemsWebsockets struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsWebsocketsID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsWebsocketsValue] `json:"value"`
}

func (r ZoneSettingEditParamsItemsWebsockets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsWebsockets) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsWebsocketsID string

const (
	ZoneSettingEditParamsItemsWebsocketsIDWebsockets ZoneSettingEditParamsItemsWebsocketsID = "websockets"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsWebsocketsEditable bool

const (
	ZoneSettingEditParamsItemsWebsocketsEditableTrue  ZoneSettingEditParamsItemsWebsocketsEditable = true
	ZoneSettingEditParamsItemsWebsocketsEditableFalse ZoneSettingEditParamsItemsWebsocketsEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsWebsocketsValue string

const (
	ZoneSettingEditParamsItemsWebsocketsValueOff ZoneSettingEditParamsItemsWebsocketsValue = "off"
	ZoneSettingEditParamsItemsWebsocketsValueOn  ZoneSettingEditParamsItemsWebsocketsValue = "on"
)
