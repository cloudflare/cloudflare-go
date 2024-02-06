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
	AdvancedDDOS                  *ZoneSettingAdvancedDDOSService
	AlwaysOnline                  *ZoneSettingAlwaysOnlineService
	AlwaysUseHTTPs                *ZoneSettingAlwaysUseHTTPService
	AutomaticHTTPsRewrites        *ZoneSettingAutomaticHTTPsRewriteService
	AutomaticPlatformOptimization *ZoneSettingAutomaticPlatformOptimizationService
	Brotli                        *ZoneSettingBrotliService
	BrowserCacheTTLs              *ZoneSettingBrowserCacheTTLService
	BrowserChecks                 *ZoneSettingBrowserCheckService
	CacheLevels                   *ZoneSettingCacheLevelService
	ChallengeTTLs                 *ZoneSettingChallengeTTLService
	Ciphers                       *ZoneSettingCipherService
	DevelopmentModes              *ZoneSettingDevelopmentModeService
	EarlyHints                    *ZoneSettingEarlyHintService
	EmailObfuscations             *ZoneSettingEmailObfuscationService
	H2Prioritizations             *ZoneSettingH2PrioritizationService
	HotlinkProtections            *ZoneSettingHotlinkProtectionService
	HTTP2s                        *ZoneSettingHTTP2Service
	HTTP3s                        *ZoneSettingHTTP3Service
	ImageResizings                *ZoneSettingImageResizingService
	IPGeolocations                *ZoneSettingIPGeolocationService
	IPV6s                         *ZoneSettingIPV6Service
	MinTLSVersions                *ZoneSettingMinTLSVersionService
	Minifies                      *ZoneSettingMinifyService
	Mirages                       *ZoneSettingMirageService
	MobileRedirects               *ZoneSettingMobileRedirectService
	NELs                          *ZoneSettingNELService
	OpportunisticEncryptions      *ZoneSettingOpportunisticEncryptionService
	OpportunisticOnions           *ZoneSettingOpportunisticOnionService
	OrangeToOranges               *ZoneSettingOrangeToOrangeService
	OriginErrorPagePassThrus      *ZoneSettingOriginErrorPagePassThrusService
	OriginMaxHTTPVersions         *ZoneSettingOriginMaxHTTPVersionService
	Polishes                      *ZoneSettingPolishService
	PrefetchPreloads              *ZoneSettingPrefetchPreloadService
	PrivacyPasses                 *ZoneSettingPrivacyPassService
	ProxyReadTimeouts             *ZoneSettingProxyReadTimeoutService
	PseudoIpv4s                   *ZoneSettingPseudoIpv4Service
	ResponseBufferings            *ZoneSettingResponseBufferingService
	RocketLoaders                 *ZoneSettingRocketLoaderService
	SecurityHeaders               *ZoneSettingSecurityHeaderService
	SecurityLevels                *ZoneSettingSecurityLevelService
	ServerSideExcludes            *ZoneSettingServerSideExcludeService
	SortQueryStringForCaches      *ZoneSettingSortQueryStringForCachService
	SSLs                          *ZoneSettingSSLService
	SSLRecommenders               *ZoneSettingSSLRecommenderService
	TLS1_3s                       *ZoneSettingTLS1_3Service
	TLSClientAuths                *ZoneSettingTLSClientAuthService
	TrueClientIPHeaders           *ZoneSettingTrueClientIPHeaderService
	WAFs                          *ZoneSettingWAFService
	Webps                         *ZoneSettingWebpService
	Websockets                    *ZoneSettingWebsocketService
}

// NewZoneSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingService(opts ...option.RequestOption) (r *ZoneSettingService) {
	r = &ZoneSettingService{}
	r.Options = opts
	r.AdvancedDDOS = NewZoneSettingAdvancedDDOSService(opts...)
	r.AlwaysOnline = NewZoneSettingAlwaysOnlineService(opts...)
	r.AlwaysUseHTTPs = NewZoneSettingAlwaysUseHTTPService(opts...)
	r.AutomaticHTTPsRewrites = NewZoneSettingAutomaticHTTPsRewriteService(opts...)
	r.AutomaticPlatformOptimization = NewZoneSettingAutomaticPlatformOptimizationService(opts...)
	r.Brotli = NewZoneSettingBrotliService(opts...)
	r.BrowserCacheTTLs = NewZoneSettingBrowserCacheTTLService(opts...)
	r.BrowserChecks = NewZoneSettingBrowserCheckService(opts...)
	r.CacheLevels = NewZoneSettingCacheLevelService(opts...)
	r.ChallengeTTLs = NewZoneSettingChallengeTTLService(opts...)
	r.Ciphers = NewZoneSettingCipherService(opts...)
	r.DevelopmentModes = NewZoneSettingDevelopmentModeService(opts...)
	r.EarlyHints = NewZoneSettingEarlyHintService(opts...)
	r.EmailObfuscations = NewZoneSettingEmailObfuscationService(opts...)
	r.H2Prioritizations = NewZoneSettingH2PrioritizationService(opts...)
	r.HotlinkProtections = NewZoneSettingHotlinkProtectionService(opts...)
	r.HTTP2s = NewZoneSettingHTTP2Service(opts...)
	r.HTTP3s = NewZoneSettingHTTP3Service(opts...)
	r.ImageResizings = NewZoneSettingImageResizingService(opts...)
	r.IPGeolocations = NewZoneSettingIPGeolocationService(opts...)
	r.IPV6s = NewZoneSettingIPV6Service(opts...)
	r.MinTLSVersions = NewZoneSettingMinTLSVersionService(opts...)
	r.Minifies = NewZoneSettingMinifyService(opts...)
	r.Mirages = NewZoneSettingMirageService(opts...)
	r.MobileRedirects = NewZoneSettingMobileRedirectService(opts...)
	r.NELs = NewZoneSettingNELService(opts...)
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
	r.SSLs = NewZoneSettingSSLService(opts...)
	r.SSLRecommenders = NewZoneSettingSSLRecommenderService(opts...)
	r.TLS1_3s = NewZoneSettingTLS1_3Service(opts...)
	r.TLSClientAuths = NewZoneSettingTLSClientAuthService(opts...)
	r.TrueClientIPHeaders = NewZoneSettingTrueClientIPHeaderService(opts...)
	r.WAFs = NewZoneSettingWAFService(opts...)
	r.Webps = NewZoneSettingWebpService(opts...)
	r.Websockets = NewZoneSettingWebsocketService(opts...)
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
// [ZoneSettingListResponseResultZonesAdvancedDDOS],
// [ZoneSettingListResponseResultZonesAlwaysOnline],
// [ZoneSettingListResponseResultZonesAlwaysUseHTTPs],
// [ZoneSettingListResponseResultZonesAutomaticHTTPsRewrites],
// [ZoneSettingListResponseResultZonesBrotli],
// [ZoneSettingListResponseResultZonesBrowserCacheTTL],
// [ZoneSettingListResponseResultZonesBrowserCheck],
// [ZoneSettingListResponseResultZonesCacheLevel],
// [ZoneSettingListResponseResultZonesChallengeTTL],
// [ZoneSettingListResponseResultZonesCiphers],
// [ZoneSettingListResponseResultZonesCnameFlattening],
// [ZoneSettingListResponseResultZonesDevelopmentMode],
// [ZoneSettingListResponseResultZonesEarlyHints],
// [ZoneSettingListResponseResultZonesEdgeCacheTTL],
// [ZoneSettingListResponseResultZonesEmailObfuscation],
// [ZoneSettingListResponseResultZonesH2Prioritization],
// [ZoneSettingListResponseResultZonesHotlinkProtection],
// [ZoneSettingListResponseResultZonesHTTP2],
// [ZoneSettingListResponseResultZonesHTTP3],
// [ZoneSettingListResponseResultZonesImageResizing],
// [ZoneSettingListResponseResultZonesIPGeolocation],
// [ZoneSettingListResponseResultZonesIPV6],
// [ZoneSettingListResponseResultZonesMaxUpload],
// [ZoneSettingListResponseResultZonesMinTLSVersion],
// [ZoneSettingListResponseResultZonesMinify],
// [ZoneSettingListResponseResultZonesMirage],
// [ZoneSettingListResponseResultZonesMobileRedirect],
// [ZoneSettingListResponseResultZonesNEL],
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
// [ZoneSettingListResponseResultZonesSSL],
// [ZoneSettingListResponseResultZonesSSLRecommender],
// [ZoneSettingListResponseResultZonesTLS1_2Only],
// [ZoneSettingListResponseResultZonesTLS1_3],
// [ZoneSettingListResponseResultZonesTLSClientAuth],
// [ZoneSettingListResponseResultZonesTrueClientIPHeader],
// [ZoneSettingListResponseResultZonesWAF],
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
type ZoneSettingListResponseResultZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesAdvancedDDOSID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingListResponseResultZonesAdvancedDDOSValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesAdvancedDDOSJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesAdvancedDDOSJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesAdvancedDDOS]
type zoneSettingListResponseResultZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesAdvancedDDOS) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesAdvancedDDOSID string

const (
	ZoneSettingListResponseResultZonesAdvancedDDOSIDAdvancedDDOS ZoneSettingListResponseResultZonesAdvancedDDOSID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesAdvancedDDOSEditable bool

const (
	ZoneSettingListResponseResultZonesAdvancedDDOSEditableTrue  ZoneSettingListResponseResultZonesAdvancedDDOSEditable = true
	ZoneSettingListResponseResultZonesAdvancedDDOSEditableFalse ZoneSettingListResponseResultZonesAdvancedDDOSEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingListResponseResultZonesAdvancedDDOSValue string

const (
	ZoneSettingListResponseResultZonesAdvancedDDOSValueOn  ZoneSettingListResponseResultZonesAdvancedDDOSValue = "on"
	ZoneSettingListResponseResultZonesAdvancedDDOSValueOff ZoneSettingListResponseResultZonesAdvancedDDOSValue = "off"
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
type ZoneSettingListResponseResultZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesBrowserCacheTTLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingListResponseResultZonesBrowserCacheTTLValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesBrowserCacheTTLJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesBrowserCacheTTLJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesBrowserCacheTTL]
type zoneSettingListResponseResultZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesBrowserCacheTTL) implementsZoneSettingListResponseResult() {
}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesBrowserCacheTTLID string

const (
	ZoneSettingListResponseResultZonesBrowserCacheTTLIDBrowserCacheTTL ZoneSettingListResponseResultZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingListResponseResultZonesBrowserCacheTTLEditableTrue  ZoneSettingListResponseResultZonesBrowserCacheTTLEditable = true
	ZoneSettingListResponseResultZonesBrowserCacheTTLEditableFalse ZoneSettingListResponseResultZonesBrowserCacheTTLEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingListResponseResultZonesBrowserCacheTTLValue float64

const (
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue0        ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 0
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue30       ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 30
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue60       ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 60
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue120      ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 120
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue300      ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 300
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue1200     ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 1200
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue1800     ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 1800
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue3600     ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 3600
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue7200     ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 7200
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue10800    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 10800
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue14400    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 14400
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue18000    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 18000
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue28800    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 28800
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue43200    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 43200
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue57600    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 57600
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue72000    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 72000
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue86400    ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 86400
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue172800   ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 172800
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue259200   ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 259200
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue345600   ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 345600
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue432000   ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 432000
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue691200   ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 691200
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue1382400  ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 1382400
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue2073600  ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 2073600
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue2678400  ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 2678400
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue5356800  ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 5356800
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue16070400 ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 16070400
	ZoneSettingListResponseResultZonesBrowserCacheTTLValue31536000 ZoneSettingListResponseResultZonesBrowserCacheTTLValue = 31536000
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
type ZoneSettingListResponseResultZonesChallengeTTL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesChallengeTTLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesChallengeTTLValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesChallengeTTLJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesChallengeTTLJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesChallengeTTL]
type zoneSettingListResponseResultZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesChallengeTTL) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesChallengeTTLID string

const (
	ZoneSettingListResponseResultZonesChallengeTTLIDChallengeTTL ZoneSettingListResponseResultZonesChallengeTTLID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesChallengeTTLEditable bool

const (
	ZoneSettingListResponseResultZonesChallengeTTLEditableTrue  ZoneSettingListResponseResultZonesChallengeTTLEditable = true
	ZoneSettingListResponseResultZonesChallengeTTLEditableFalse ZoneSettingListResponseResultZonesChallengeTTLEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesChallengeTTLValue float64

const (
	ZoneSettingListResponseResultZonesChallengeTTLValue300      ZoneSettingListResponseResultZonesChallengeTTLValue = 300
	ZoneSettingListResponseResultZonesChallengeTTLValue900      ZoneSettingListResponseResultZonesChallengeTTLValue = 900
	ZoneSettingListResponseResultZonesChallengeTTLValue1800     ZoneSettingListResponseResultZonesChallengeTTLValue = 1800
	ZoneSettingListResponseResultZonesChallengeTTLValue2700     ZoneSettingListResponseResultZonesChallengeTTLValue = 2700
	ZoneSettingListResponseResultZonesChallengeTTLValue3600     ZoneSettingListResponseResultZonesChallengeTTLValue = 3600
	ZoneSettingListResponseResultZonesChallengeTTLValue7200     ZoneSettingListResponseResultZonesChallengeTTLValue = 7200
	ZoneSettingListResponseResultZonesChallengeTTLValue10800    ZoneSettingListResponseResultZonesChallengeTTLValue = 10800
	ZoneSettingListResponseResultZonesChallengeTTLValue14400    ZoneSettingListResponseResultZonesChallengeTTLValue = 14400
	ZoneSettingListResponseResultZonesChallengeTTLValue28800    ZoneSettingListResponseResultZonesChallengeTTLValue = 28800
	ZoneSettingListResponseResultZonesChallengeTTLValue57600    ZoneSettingListResponseResultZonesChallengeTTLValue = 57600
	ZoneSettingListResponseResultZonesChallengeTTLValue86400    ZoneSettingListResponseResultZonesChallengeTTLValue = 86400
	ZoneSettingListResponseResultZonesChallengeTTLValue604800   ZoneSettingListResponseResultZonesChallengeTTLValue = 604800
	ZoneSettingListResponseResultZonesChallengeTTLValue2592000  ZoneSettingListResponseResultZonesChallengeTTLValue = 2592000
	ZoneSettingListResponseResultZonesChallengeTTLValue31536000 ZoneSettingListResponseResultZonesChallengeTTLValue = 31536000
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
type ZoneSettingListResponseResultZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesEdgeCacheTTLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value ZoneSettingListResponseResultZonesEdgeCacheTTLValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesEdgeCacheTTLJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesEdgeCacheTTLJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesEdgeCacheTTL]
type zoneSettingListResponseResultZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesEdgeCacheTTL) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesEdgeCacheTTLID string

const (
	ZoneSettingListResponseResultZonesEdgeCacheTTLIDEdgeCacheTTL ZoneSettingListResponseResultZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingListResponseResultZonesEdgeCacheTTLEditableTrue  ZoneSettingListResponseResultZonesEdgeCacheTTLEditable = true
	ZoneSettingListResponseResultZonesEdgeCacheTTLEditableFalse ZoneSettingListResponseResultZonesEdgeCacheTTLEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingListResponseResultZonesEdgeCacheTTLValue float64

const (
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue30     ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 30
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue60     ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 60
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue300    ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 300
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue1200   ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 1200
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue1800   ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 1800
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue3600   ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 3600
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue7200   ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 7200
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue10800  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 10800
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue14400  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 14400
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue18000  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 18000
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue28800  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 28800
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue43200  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 43200
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue57600  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 57600
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue72000  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 72000
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue86400  ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 86400
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue172800 ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 172800
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue259200 ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 259200
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue345600 ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 345600
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue432000 ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 432000
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue518400 ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 518400
	ZoneSettingListResponseResultZonesEdgeCacheTTLValue604800 ZoneSettingListResponseResultZonesEdgeCacheTTLValue = 604800
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
type ZoneSettingListResponseResultZonesHTTP2 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesHTTP2ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingListResponseResultZonesHTTP2Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesHTTP2JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesHTTP2JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesHTTP2]
type zoneSettingListResponseResultZonesHTTP2JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesHTTP2) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesHTTP2ID string

const (
	ZoneSettingListResponseResultZonesHTTP2IDHTTP2 ZoneSettingListResponseResultZonesHTTP2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesHTTP2Editable bool

const (
	ZoneSettingListResponseResultZonesHTTP2EditableTrue  ZoneSettingListResponseResultZonesHTTP2Editable = true
	ZoneSettingListResponseResultZonesHTTP2EditableFalse ZoneSettingListResponseResultZonesHTTP2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingListResponseResultZonesHTTP2Value string

const (
	ZoneSettingListResponseResultZonesHTTP2ValueOn  ZoneSettingListResponseResultZonesHTTP2Value = "on"
	ZoneSettingListResponseResultZonesHTTP2ValueOff ZoneSettingListResponseResultZonesHTTP2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingListResponseResultZonesHTTP3 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesHTTP3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingListResponseResultZonesHTTP3Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesHTTP3JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesHTTP3JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesHTTP3]
type zoneSettingListResponseResultZonesHTTP3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesHTTP3) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesHTTP3ID string

const (
	ZoneSettingListResponseResultZonesHTTP3IDHTTP3 ZoneSettingListResponseResultZonesHTTP3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesHTTP3Editable bool

const (
	ZoneSettingListResponseResultZonesHTTP3EditableTrue  ZoneSettingListResponseResultZonesHTTP3Editable = true
	ZoneSettingListResponseResultZonesHTTP3EditableFalse ZoneSettingListResponseResultZonesHTTP3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingListResponseResultZonesHTTP3Value string

const (
	ZoneSettingListResponseResultZonesHTTP3ValueOn  ZoneSettingListResponseResultZonesHTTP3Value = "on"
	ZoneSettingListResponseResultZonesHTTP3ValueOff ZoneSettingListResponseResultZonesHTTP3Value = "off"
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
type ZoneSettingListResponseResultZonesIPV6 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesIPV6ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesIPV6Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesIPV6JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesIPV6JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesIPV6]
type zoneSettingListResponseResultZonesIPV6JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesIPV6) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesIPV6ID string

const (
	ZoneSettingListResponseResultZonesIPV6IDIPV6 ZoneSettingListResponseResultZonesIPV6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesIPV6Editable bool

const (
	ZoneSettingListResponseResultZonesIPV6EditableTrue  ZoneSettingListResponseResultZonesIPV6Editable = true
	ZoneSettingListResponseResultZonesIPV6EditableFalse ZoneSettingListResponseResultZonesIPV6Editable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesIPV6Value string

const (
	ZoneSettingListResponseResultZonesIPV6ValueOff ZoneSettingListResponseResultZonesIPV6Value = "off"
	ZoneSettingListResponseResultZonesIPV6ValueOn  ZoneSettingListResponseResultZonesIPV6Value = "on"
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
type ZoneSettingListResponseResultZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesMinTLSVersionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesMinTLSVersionValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesMinTLSVersionJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesMinTLSVersionJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesMinTLSVersion]
type zoneSettingListResponseResultZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesMinTLSVersion) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesMinTLSVersionID string

const (
	ZoneSettingListResponseResultZonesMinTLSVersionIDMinTLSVersion ZoneSettingListResponseResultZonesMinTLSVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesMinTLSVersionEditable bool

const (
	ZoneSettingListResponseResultZonesMinTLSVersionEditableTrue  ZoneSettingListResponseResultZonesMinTLSVersionEditable = true
	ZoneSettingListResponseResultZonesMinTLSVersionEditableFalse ZoneSettingListResponseResultZonesMinTLSVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesMinTLSVersionValue string

const (
	ZoneSettingListResponseResultZonesMinTLSVersionValue1_0 ZoneSettingListResponseResultZonesMinTLSVersionValue = "1.0"
	ZoneSettingListResponseResultZonesMinTLSVersionValue1_1 ZoneSettingListResponseResultZonesMinTLSVersionValue = "1.1"
	ZoneSettingListResponseResultZonesMinTLSVersionValue1_2 ZoneSettingListResponseResultZonesMinTLSVersionValue = "1.2"
	ZoneSettingListResponseResultZonesMinTLSVersionValue1_3 ZoneSettingListResponseResultZonesMinTLSVersionValue = "1.3"
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
type ZoneSettingListResponseResultZonesNEL struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseResultZonesNELID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesNELValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesNELJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesNELJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesNEL]
type zoneSettingListResponseResultZonesNELJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesNEL) implementsZoneSettingListResponseResult() {}

// Zone setting identifier.
type ZoneSettingListResponseResultZonesNELID string

const (
	ZoneSettingListResponseResultZonesNELIDNEL ZoneSettingListResponseResultZonesNELID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesNELEditable bool

const (
	ZoneSettingListResponseResultZonesNELEditableTrue  ZoneSettingListResponseResultZonesNELEditable = true
	ZoneSettingListResponseResultZonesNELEditableFalse ZoneSettingListResponseResultZonesNELEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesNELValue struct {
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSettingListResponseResultZonesNELValueJSON `json:"-"`
}

// zoneSettingListResponseResultZonesNELValueJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesNELValue]
type zoneSettingListResponseResultZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesNELValue) UnmarshalJSON(data []byte) (err error) {
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
type ZoneSettingListResponseResultZonesSSL struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesSSLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingListResponseResultZonesSSLValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesSSLJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesSSLJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesSSL]
type zoneSettingListResponseResultZonesSSLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSSL) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesSSLID string

const (
	ZoneSettingListResponseResultZonesSSLIDSSL ZoneSettingListResponseResultZonesSSLID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesSSLEditable bool

const (
	ZoneSettingListResponseResultZonesSSLEditableTrue  ZoneSettingListResponseResultZonesSSLEditable = true
	ZoneSettingListResponseResultZonesSSLEditableFalse ZoneSettingListResponseResultZonesSSLEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingListResponseResultZonesSSLValue string

const (
	ZoneSettingListResponseResultZonesSSLValueOff      ZoneSettingListResponseResultZonesSSLValue = "off"
	ZoneSettingListResponseResultZonesSSLValueFlexible ZoneSettingListResponseResultZonesSSLValue = "flexible"
	ZoneSettingListResponseResultZonesSSLValueFull     ZoneSettingListResponseResultZonesSSLValue = "full"
	ZoneSettingListResponseResultZonesSSLValueStrict   ZoneSettingListResponseResultZonesSSLValue = "strict"
)

type ZoneSettingListResponseResultZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingListResponseResultZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                                 `json:"enabled"`
	JSON    zoneSettingListResponseResultZonesSSLRecommenderJSON `json:"-"`
}

// zoneSettingListResponseResultZonesSSLRecommenderJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesSSLRecommender]
type zoneSettingListResponseResultZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesSSLRecommender) implementsZoneSettingListResponseResult() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingListResponseResultZonesSSLRecommenderID string

const (
	ZoneSettingListResponseResultZonesSSLRecommenderIDSSLRecommender ZoneSettingListResponseResultZonesSSLRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type ZoneSettingListResponseResultZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingListResponseResultZonesTLS1_2OnlyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesTLS1_2OnlyValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesTls1_2OnlyJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTls1_2OnlyJSON contains the JSON metadata for
// the struct [ZoneSettingListResponseResultZonesTLS1_2Only]
type zoneSettingListResponseResultZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTLS1_2Only) implementsZoneSettingListResponseResult() {}

// Zone setting identifier.
type ZoneSettingListResponseResultZonesTLS1_2OnlyID string

const (
	ZoneSettingListResponseResultZonesTLS1_2OnlyIDTLS1_2Only ZoneSettingListResponseResultZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingListResponseResultZonesTLS1_2OnlyEditableTrue  ZoneSettingListResponseResultZonesTLS1_2OnlyEditable = true
	ZoneSettingListResponseResultZonesTLS1_2OnlyEditableFalse ZoneSettingListResponseResultZonesTLS1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesTLS1_2OnlyValue string

const (
	ZoneSettingListResponseResultZonesTLS1_2OnlyValueOff ZoneSettingListResponseResultZonesTLS1_2OnlyValue = "off"
	ZoneSettingListResponseResultZonesTLS1_2OnlyValueOn  ZoneSettingListResponseResultZonesTLS1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingListResponseResultZonesTLS1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesTLS1_3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingListResponseResultZonesTLS1_3Value `json:"value"`
	JSON  zoneSettingListResponseResultZonesTls1_3JSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTls1_3JSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesTLS1_3]
type zoneSettingListResponseResultZonesTls1_3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTLS1_3) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesTLS1_3ID string

const (
	ZoneSettingListResponseResultZonesTLS1_3IDTLS1_3 ZoneSettingListResponseResultZonesTLS1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTLS1_3Editable bool

const (
	ZoneSettingListResponseResultZonesTLS1_3EditableTrue  ZoneSettingListResponseResultZonesTLS1_3Editable = true
	ZoneSettingListResponseResultZonesTLS1_3EditableFalse ZoneSettingListResponseResultZonesTLS1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingListResponseResultZonesTLS1_3Value string

const (
	ZoneSettingListResponseResultZonesTLS1_3ValueOn  ZoneSettingListResponseResultZonesTLS1_3Value = "on"
	ZoneSettingListResponseResultZonesTLS1_3ValueOff ZoneSettingListResponseResultZonesTLS1_3Value = "off"
	ZoneSettingListResponseResultZonesTLS1_3ValueZrt ZoneSettingListResponseResultZonesTLS1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingListResponseResultZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesTLSClientAuthID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingListResponseResultZonesTLSClientAuthValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesTLSClientAuthJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesTLSClientAuthJSON contains the JSON metadata
// for the struct [ZoneSettingListResponseResultZonesTLSClientAuth]
type zoneSettingListResponseResultZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesTLSClientAuth) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesTLSClientAuthID string

const (
	ZoneSettingListResponseResultZonesTLSClientAuthIDTLSClientAuth ZoneSettingListResponseResultZonesTLSClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesTLSClientAuthEditable bool

const (
	ZoneSettingListResponseResultZonesTLSClientAuthEditableTrue  ZoneSettingListResponseResultZonesTLSClientAuthEditable = true
	ZoneSettingListResponseResultZonesTLSClientAuthEditableFalse ZoneSettingListResponseResultZonesTLSClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingListResponseResultZonesTLSClientAuthValue string

const (
	ZoneSettingListResponseResultZonesTLSClientAuthValueOn  ZoneSettingListResponseResultZonesTLSClientAuthValue = "on"
	ZoneSettingListResponseResultZonesTLSClientAuthValueOff ZoneSettingListResponseResultZonesTLSClientAuthValue = "off"
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
type ZoneSettingListResponseResultZonesWAF struct {
	// ID of the zone setting.
	ID ZoneSettingListResponseResultZonesWAFID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingListResponseResultZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingListResponseResultZonesWAFValue `json:"value"`
	JSON  zoneSettingListResponseResultZonesWAFJSON  `json:"-"`
}

// zoneSettingListResponseResultZonesWAFJSON contains the JSON metadata for the
// struct [ZoneSettingListResponseResultZonesWAF]
type zoneSettingListResponseResultZonesWAFJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingListResponseResultZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingListResponseResultZonesWAF) implementsZoneSettingListResponseResult() {}

// ID of the zone setting.
type ZoneSettingListResponseResultZonesWAFID string

const (
	ZoneSettingListResponseResultZonesWAFIDWAF ZoneSettingListResponseResultZonesWAFID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingListResponseResultZonesWAFEditable bool

const (
	ZoneSettingListResponseResultZonesWAFEditableTrue  ZoneSettingListResponseResultZonesWAFEditable = true
	ZoneSettingListResponseResultZonesWAFEditableFalse ZoneSettingListResponseResultZonesWAFEditable = false
)

// Value of the zone setting.
type ZoneSettingListResponseResultZonesWAFValue string

const (
	ZoneSettingListResponseResultZonesWAFValueOn  ZoneSettingListResponseResultZonesWAFValue = "on"
	ZoneSettingListResponseResultZonesWAFValueOff ZoneSettingListResponseResultZonesWAFValue = "off"
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
// [ZoneSettingEditResponseResultZonesAdvancedDDOS],
// [ZoneSettingEditResponseResultZonesAlwaysOnline],
// [ZoneSettingEditResponseResultZonesAlwaysUseHTTPs],
// [ZoneSettingEditResponseResultZonesAutomaticHTTPsRewrites],
// [ZoneSettingEditResponseResultZonesBrotli],
// [ZoneSettingEditResponseResultZonesBrowserCacheTTL],
// [ZoneSettingEditResponseResultZonesBrowserCheck],
// [ZoneSettingEditResponseResultZonesCacheLevel],
// [ZoneSettingEditResponseResultZonesChallengeTTL],
// [ZoneSettingEditResponseResultZonesCiphers],
// [ZoneSettingEditResponseResultZonesCnameFlattening],
// [ZoneSettingEditResponseResultZonesDevelopmentMode],
// [ZoneSettingEditResponseResultZonesEarlyHints],
// [ZoneSettingEditResponseResultZonesEdgeCacheTTL],
// [ZoneSettingEditResponseResultZonesEmailObfuscation],
// [ZoneSettingEditResponseResultZonesH2Prioritization],
// [ZoneSettingEditResponseResultZonesHotlinkProtection],
// [ZoneSettingEditResponseResultZonesHTTP2],
// [ZoneSettingEditResponseResultZonesHTTP3],
// [ZoneSettingEditResponseResultZonesImageResizing],
// [ZoneSettingEditResponseResultZonesIPGeolocation],
// [ZoneSettingEditResponseResultZonesIPV6],
// [ZoneSettingEditResponseResultZonesMaxUpload],
// [ZoneSettingEditResponseResultZonesMinTLSVersion],
// [ZoneSettingEditResponseResultZonesMinify],
// [ZoneSettingEditResponseResultZonesMirage],
// [ZoneSettingEditResponseResultZonesMobileRedirect],
// [ZoneSettingEditResponseResultZonesNEL],
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
// [ZoneSettingEditResponseResultZonesSSL],
// [ZoneSettingEditResponseResultZonesSSLRecommender],
// [ZoneSettingEditResponseResultZonesTLS1_2Only],
// [ZoneSettingEditResponseResultZonesTLS1_3],
// [ZoneSettingEditResponseResultZonesTLSClientAuth],
// [ZoneSettingEditResponseResultZonesTrueClientIPHeader],
// [ZoneSettingEditResponseResultZonesWAF],
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
type ZoneSettingEditResponseResultZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesAdvancedDDOSID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesAdvancedDDOSEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value ZoneSettingEditResponseResultZonesAdvancedDDOSValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesAdvancedDDOSJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesAdvancedDDOSJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesAdvancedDDOS]
type zoneSettingEditResponseResultZonesAdvancedDDOSJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesAdvancedDDOS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesAdvancedDDOS) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesAdvancedDDOSID string

const (
	ZoneSettingEditResponseResultZonesAdvancedDDOSIDAdvancedDDOS ZoneSettingEditResponseResultZonesAdvancedDDOSID = "advanced_ddos"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesAdvancedDDOSEditable bool

const (
	ZoneSettingEditResponseResultZonesAdvancedDDOSEditableTrue  ZoneSettingEditResponseResultZonesAdvancedDDOSEditable = true
	ZoneSettingEditResponseResultZonesAdvancedDDOSEditableFalse ZoneSettingEditResponseResultZonesAdvancedDDOSEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingEditResponseResultZonesAdvancedDDOSValue string

const (
	ZoneSettingEditResponseResultZonesAdvancedDDOSValueOn  ZoneSettingEditResponseResultZonesAdvancedDDOSValue = "on"
	ZoneSettingEditResponseResultZonesAdvancedDDOSValueOff ZoneSettingEditResponseResultZonesAdvancedDDOSValue = "off"
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
type ZoneSettingEditResponseResultZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesBrowserCacheTTLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesBrowserCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value ZoneSettingEditResponseResultZonesBrowserCacheTTLValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesBrowserCacheTTLJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesBrowserCacheTTLJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesBrowserCacheTTL]
type zoneSettingEditResponseResultZonesBrowserCacheTTLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesBrowserCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesBrowserCacheTTL) implementsZoneSettingEditResponseResult() {
}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesBrowserCacheTTLID string

const (
	ZoneSettingEditResponseResultZonesBrowserCacheTTLIDBrowserCacheTTL ZoneSettingEditResponseResultZonesBrowserCacheTTLID = "browser_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingEditResponseResultZonesBrowserCacheTTLEditableTrue  ZoneSettingEditResponseResultZonesBrowserCacheTTLEditable = true
	ZoneSettingEditResponseResultZonesBrowserCacheTTLEditableFalse ZoneSettingEditResponseResultZonesBrowserCacheTTLEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
type ZoneSettingEditResponseResultZonesBrowserCacheTTLValue float64

const (
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue0        ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 0
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue30       ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 30
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue60       ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 60
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue120      ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 120
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue300      ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 300
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue1200     ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 1200
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue1800     ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 1800
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue3600     ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 3600
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue7200     ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 7200
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue10800    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 10800
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue14400    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 14400
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue18000    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 18000
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue28800    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 28800
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue43200    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 43200
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue57600    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 57600
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue72000    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 72000
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue86400    ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 86400
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue172800   ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 172800
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue259200   ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 259200
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue345600   ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 345600
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue432000   ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 432000
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue691200   ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 691200
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue1382400  ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 1382400
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue2073600  ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 2073600
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue2678400  ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 2678400
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue5356800  ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 5356800
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue16070400 ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 16070400
	ZoneSettingEditResponseResultZonesBrowserCacheTTLValue31536000 ZoneSettingEditResponseResultZonesBrowserCacheTTLValue = 31536000
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
type ZoneSettingEditResponseResultZonesChallengeTTL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesChallengeTTLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesChallengeTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesChallengeTTLValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesChallengeTTLJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesChallengeTTLJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesChallengeTTL]
type zoneSettingEditResponseResultZonesChallengeTTLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesChallengeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesChallengeTTL) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesChallengeTTLID string

const (
	ZoneSettingEditResponseResultZonesChallengeTTLIDChallengeTTL ZoneSettingEditResponseResultZonesChallengeTTLID = "challenge_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesChallengeTTLEditable bool

const (
	ZoneSettingEditResponseResultZonesChallengeTTLEditableTrue  ZoneSettingEditResponseResultZonesChallengeTTLEditable = true
	ZoneSettingEditResponseResultZonesChallengeTTLEditableFalse ZoneSettingEditResponseResultZonesChallengeTTLEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesChallengeTTLValue float64

const (
	ZoneSettingEditResponseResultZonesChallengeTTLValue300      ZoneSettingEditResponseResultZonesChallengeTTLValue = 300
	ZoneSettingEditResponseResultZonesChallengeTTLValue900      ZoneSettingEditResponseResultZonesChallengeTTLValue = 900
	ZoneSettingEditResponseResultZonesChallengeTTLValue1800     ZoneSettingEditResponseResultZonesChallengeTTLValue = 1800
	ZoneSettingEditResponseResultZonesChallengeTTLValue2700     ZoneSettingEditResponseResultZonesChallengeTTLValue = 2700
	ZoneSettingEditResponseResultZonesChallengeTTLValue3600     ZoneSettingEditResponseResultZonesChallengeTTLValue = 3600
	ZoneSettingEditResponseResultZonesChallengeTTLValue7200     ZoneSettingEditResponseResultZonesChallengeTTLValue = 7200
	ZoneSettingEditResponseResultZonesChallengeTTLValue10800    ZoneSettingEditResponseResultZonesChallengeTTLValue = 10800
	ZoneSettingEditResponseResultZonesChallengeTTLValue14400    ZoneSettingEditResponseResultZonesChallengeTTLValue = 14400
	ZoneSettingEditResponseResultZonesChallengeTTLValue28800    ZoneSettingEditResponseResultZonesChallengeTTLValue = 28800
	ZoneSettingEditResponseResultZonesChallengeTTLValue57600    ZoneSettingEditResponseResultZonesChallengeTTLValue = 57600
	ZoneSettingEditResponseResultZonesChallengeTTLValue86400    ZoneSettingEditResponseResultZonesChallengeTTLValue = 86400
	ZoneSettingEditResponseResultZonesChallengeTTLValue604800   ZoneSettingEditResponseResultZonesChallengeTTLValue = 604800
	ZoneSettingEditResponseResultZonesChallengeTTLValue2592000  ZoneSettingEditResponseResultZonesChallengeTTLValue = 2592000
	ZoneSettingEditResponseResultZonesChallengeTTLValue31536000 ZoneSettingEditResponseResultZonesChallengeTTLValue = 31536000
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
type ZoneSettingEditResponseResultZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesEdgeCacheTTLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesEdgeCacheTTLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value ZoneSettingEditResponseResultZonesEdgeCacheTTLValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesEdgeCacheTTLJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesEdgeCacheTTLJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesEdgeCacheTTL]
type zoneSettingEditResponseResultZonesEdgeCacheTTLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesEdgeCacheTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesEdgeCacheTTL) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesEdgeCacheTTLID string

const (
	ZoneSettingEditResponseResultZonesEdgeCacheTTLIDEdgeCacheTTL ZoneSettingEditResponseResultZonesEdgeCacheTTLID = "edge_cache_ttl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingEditResponseResultZonesEdgeCacheTTLEditableTrue  ZoneSettingEditResponseResultZonesEdgeCacheTTLEditable = true
	ZoneSettingEditResponseResultZonesEdgeCacheTTLEditableFalse ZoneSettingEditResponseResultZonesEdgeCacheTTLEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
type ZoneSettingEditResponseResultZonesEdgeCacheTTLValue float64

const (
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue30     ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 30
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue60     ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 60
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue300    ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 300
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue1200   ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 1200
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue1800   ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 1800
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue3600   ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 3600
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue7200   ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 7200
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue10800  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 10800
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue14400  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 14400
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue18000  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 18000
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue28800  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 28800
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue43200  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 43200
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue57600  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 57600
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue72000  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 72000
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue86400  ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 86400
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue172800 ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 172800
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue259200 ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 259200
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue345600 ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 345600
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue432000 ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 432000
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue518400 ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 518400
	ZoneSettingEditResponseResultZonesEdgeCacheTTLValue604800 ZoneSettingEditResponseResultZonesEdgeCacheTTLValue = 604800
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
type ZoneSettingEditResponseResultZonesHTTP2 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesHTTP2ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesHTTP2Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP2 setting.
	Value ZoneSettingEditResponseResultZonesHTTP2Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesHTTP2JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesHTTP2JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesHTTP2]
type zoneSettingEditResponseResultZonesHTTP2JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesHTTP2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesHTTP2) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesHTTP2ID string

const (
	ZoneSettingEditResponseResultZonesHTTP2IDHTTP2 ZoneSettingEditResponseResultZonesHTTP2ID = "http2"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesHTTP2Editable bool

const (
	ZoneSettingEditResponseResultZonesHTTP2EditableTrue  ZoneSettingEditResponseResultZonesHTTP2Editable = true
	ZoneSettingEditResponseResultZonesHTTP2EditableFalse ZoneSettingEditResponseResultZonesHTTP2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingEditResponseResultZonesHTTP2Value string

const (
	ZoneSettingEditResponseResultZonesHTTP2ValueOn  ZoneSettingEditResponseResultZonesHTTP2Value = "on"
	ZoneSettingEditResponseResultZonesHTTP2ValueOff ZoneSettingEditResponseResultZonesHTTP2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingEditResponseResultZonesHTTP3 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesHTTP3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesHTTP3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the HTTP3 setting.
	Value ZoneSettingEditResponseResultZonesHTTP3Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesHTTP3JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesHTTP3JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesHTTP3]
type zoneSettingEditResponseResultZonesHTTP3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesHTTP3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesHTTP3) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesHTTP3ID string

const (
	ZoneSettingEditResponseResultZonesHTTP3IDHTTP3 ZoneSettingEditResponseResultZonesHTTP3ID = "http3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesHTTP3Editable bool

const (
	ZoneSettingEditResponseResultZonesHTTP3EditableTrue  ZoneSettingEditResponseResultZonesHTTP3Editable = true
	ZoneSettingEditResponseResultZonesHTTP3EditableFalse ZoneSettingEditResponseResultZonesHTTP3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingEditResponseResultZonesHTTP3Value string

const (
	ZoneSettingEditResponseResultZonesHTTP3ValueOn  ZoneSettingEditResponseResultZonesHTTP3Value = "on"
	ZoneSettingEditResponseResultZonesHTTP3ValueOff ZoneSettingEditResponseResultZonesHTTP3Value = "off"
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
type ZoneSettingEditResponseResultZonesIPV6 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesIPV6ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesIPV6Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesIPV6Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesIPV6JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesIPV6JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesIPV6]
type zoneSettingEditResponseResultZonesIPV6JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesIPV6) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesIPV6ID string

const (
	ZoneSettingEditResponseResultZonesIPV6IDIPV6 ZoneSettingEditResponseResultZonesIPV6ID = "ipv6"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesIPV6Editable bool

const (
	ZoneSettingEditResponseResultZonesIPV6EditableTrue  ZoneSettingEditResponseResultZonesIPV6Editable = true
	ZoneSettingEditResponseResultZonesIPV6EditableFalse ZoneSettingEditResponseResultZonesIPV6Editable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesIPV6Value string

const (
	ZoneSettingEditResponseResultZonesIPV6ValueOff ZoneSettingEditResponseResultZonesIPV6Value = "off"
	ZoneSettingEditResponseResultZonesIPV6ValueOn  ZoneSettingEditResponseResultZonesIPV6Value = "on"
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
type ZoneSettingEditResponseResultZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesMinTLSVersionID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesMinTLSVersionEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesMinTLSVersionValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesMinTLSVersionJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesMinTLSVersionJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesMinTLSVersion]
type zoneSettingEditResponseResultZonesMinTLSVersionJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesMinTLSVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesMinTLSVersion) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesMinTLSVersionID string

const (
	ZoneSettingEditResponseResultZonesMinTLSVersionIDMinTLSVersion ZoneSettingEditResponseResultZonesMinTLSVersionID = "min_tls_version"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesMinTLSVersionEditable bool

const (
	ZoneSettingEditResponseResultZonesMinTLSVersionEditableTrue  ZoneSettingEditResponseResultZonesMinTLSVersionEditable = true
	ZoneSettingEditResponseResultZonesMinTLSVersionEditableFalse ZoneSettingEditResponseResultZonesMinTLSVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesMinTLSVersionValue string

const (
	ZoneSettingEditResponseResultZonesMinTLSVersionValue1_0 ZoneSettingEditResponseResultZonesMinTLSVersionValue = "1.0"
	ZoneSettingEditResponseResultZonesMinTLSVersionValue1_1 ZoneSettingEditResponseResultZonesMinTLSVersionValue = "1.1"
	ZoneSettingEditResponseResultZonesMinTLSVersionValue1_2 ZoneSettingEditResponseResultZonesMinTLSVersionValue = "1.2"
	ZoneSettingEditResponseResultZonesMinTLSVersionValue1_3 ZoneSettingEditResponseResultZonesMinTLSVersionValue = "1.3"
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
type ZoneSettingEditResponseResultZonesNEL struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseResultZonesNELID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesNELEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesNELValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesNELJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesNELJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesNEL]
type zoneSettingEditResponseResultZonesNELJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesNEL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesNEL) implementsZoneSettingEditResponseResult() {}

// Zone setting identifier.
type ZoneSettingEditResponseResultZonesNELID string

const (
	ZoneSettingEditResponseResultZonesNELIDNEL ZoneSettingEditResponseResultZonesNELID = "nel"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesNELEditable bool

const (
	ZoneSettingEditResponseResultZonesNELEditableTrue  ZoneSettingEditResponseResultZonesNELEditable = true
	ZoneSettingEditResponseResultZonesNELEditableFalse ZoneSettingEditResponseResultZonesNELEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesNELValue struct {
	Enabled bool                                           `json:"enabled"`
	JSON    zoneSettingEditResponseResultZonesNELValueJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesNELValueJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesNELValue]
type zoneSettingEditResponseResultZonesNELValueJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesNELValue) UnmarshalJSON(data []byte) (err error) {
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
type ZoneSettingEditResponseResultZonesSSL struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesSSLID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesSSLEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value ZoneSettingEditResponseResultZonesSSLValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesSSLJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesSSLJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesSSL]
type zoneSettingEditResponseResultZonesSSLJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSSL) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesSSLID string

const (
	ZoneSettingEditResponseResultZonesSSLIDSSL ZoneSettingEditResponseResultZonesSSLID = "ssl"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesSSLEditable bool

const (
	ZoneSettingEditResponseResultZonesSSLEditableTrue  ZoneSettingEditResponseResultZonesSSLEditable = true
	ZoneSettingEditResponseResultZonesSSLEditableFalse ZoneSettingEditResponseResultZonesSSLEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingEditResponseResultZonesSSLValue string

const (
	ZoneSettingEditResponseResultZonesSSLValueOff      ZoneSettingEditResponseResultZonesSSLValue = "off"
	ZoneSettingEditResponseResultZonesSSLValueFlexible ZoneSettingEditResponseResultZonesSSLValue = "flexible"
	ZoneSettingEditResponseResultZonesSSLValueFull     ZoneSettingEditResponseResultZonesSSLValue = "full"
	ZoneSettingEditResponseResultZonesSSLValueStrict   ZoneSettingEditResponseResultZonesSSLValue = "strict"
)

type ZoneSettingEditResponseResultZonesSSLRecommender struct {
	// Enrollment value for SSL/TLS Recommender.
	ID ZoneSettingEditResponseResultZonesSSLRecommenderID `json:"id"`
	// ssl-recommender enrollment setting.
	Enabled bool                                                 `json:"enabled"`
	JSON    zoneSettingEditResponseResultZonesSSLRecommenderJSON `json:"-"`
}

// zoneSettingEditResponseResultZonesSSLRecommenderJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesSSLRecommender]
type zoneSettingEditResponseResultZonesSSLRecommenderJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesSSLRecommender) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesSSLRecommender) implementsZoneSettingEditResponseResult() {}

// Enrollment value for SSL/TLS Recommender.
type ZoneSettingEditResponseResultZonesSSLRecommenderID string

const (
	ZoneSettingEditResponseResultZonesSSLRecommenderIDSSLRecommender ZoneSettingEditResponseResultZonesSSLRecommenderID = "ssl_recommender"
)

// Only allows TLS1.2.
type ZoneSettingEditResponseResultZonesTLS1_2Only struct {
	// Zone setting identifier.
	ID ZoneSettingEditResponseResultZonesTLS1_2OnlyID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTLS1_2OnlyEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesTLS1_2OnlyValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTls1_2OnlyJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTls1_2OnlyJSON contains the JSON metadata for
// the struct [ZoneSettingEditResponseResultZonesTLS1_2Only]
type zoneSettingEditResponseResultZonesTls1_2OnlyJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTLS1_2Only) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTLS1_2Only) implementsZoneSettingEditResponseResult() {}

// Zone setting identifier.
type ZoneSettingEditResponseResultZonesTLS1_2OnlyID string

const (
	ZoneSettingEditResponseResultZonesTLS1_2OnlyIDTLS1_2Only ZoneSettingEditResponseResultZonesTLS1_2OnlyID = "tls_1_2_only"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingEditResponseResultZonesTLS1_2OnlyEditableTrue  ZoneSettingEditResponseResultZonesTLS1_2OnlyEditable = true
	ZoneSettingEditResponseResultZonesTLS1_2OnlyEditableFalse ZoneSettingEditResponseResultZonesTLS1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesTLS1_2OnlyValue string

const (
	ZoneSettingEditResponseResultZonesTLS1_2OnlyValueOff ZoneSettingEditResponseResultZonesTLS1_2OnlyValue = "off"
	ZoneSettingEditResponseResultZonesTLS1_2OnlyValueOn  ZoneSettingEditResponseResultZonesTLS1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditResponseResultZonesTLS1_3 struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesTLS1_3ID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTLS1_3Editable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value ZoneSettingEditResponseResultZonesTLS1_3Value `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTls1_3JSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTls1_3JSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesTLS1_3]
type zoneSettingEditResponseResultZonesTls1_3JSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTLS1_3) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTLS1_3) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesTLS1_3ID string

const (
	ZoneSettingEditResponseResultZonesTLS1_3IDTLS1_3 ZoneSettingEditResponseResultZonesTLS1_3ID = "tls_1_3"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTLS1_3Editable bool

const (
	ZoneSettingEditResponseResultZonesTLS1_3EditableTrue  ZoneSettingEditResponseResultZonesTLS1_3Editable = true
	ZoneSettingEditResponseResultZonesTLS1_3EditableFalse ZoneSettingEditResponseResultZonesTLS1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditResponseResultZonesTLS1_3Value string

const (
	ZoneSettingEditResponseResultZonesTLS1_3ValueOn  ZoneSettingEditResponseResultZonesTLS1_3Value = "on"
	ZoneSettingEditResponseResultZonesTLS1_3ValueOff ZoneSettingEditResponseResultZonesTLS1_3Value = "off"
	ZoneSettingEditResponseResultZonesTLS1_3ValueZrt ZoneSettingEditResponseResultZonesTLS1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditResponseResultZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesTLSClientAuthID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesTLSClientAuthEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// value of the zone setting.
	Value ZoneSettingEditResponseResultZonesTLSClientAuthValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesTLSClientAuthJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesTLSClientAuthJSON contains the JSON metadata
// for the struct [ZoneSettingEditResponseResultZonesTLSClientAuth]
type zoneSettingEditResponseResultZonesTLSClientAuthJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesTLSClientAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesTLSClientAuth) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesTLSClientAuthID string

const (
	ZoneSettingEditResponseResultZonesTLSClientAuthIDTLSClientAuth ZoneSettingEditResponseResultZonesTLSClientAuthID = "tls_client_auth"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesTLSClientAuthEditable bool

const (
	ZoneSettingEditResponseResultZonesTLSClientAuthEditableTrue  ZoneSettingEditResponseResultZonesTLSClientAuthEditable = true
	ZoneSettingEditResponseResultZonesTLSClientAuthEditableFalse ZoneSettingEditResponseResultZonesTLSClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingEditResponseResultZonesTLSClientAuthValue string

const (
	ZoneSettingEditResponseResultZonesTLSClientAuthValueOn  ZoneSettingEditResponseResultZonesTLSClientAuthValue = "on"
	ZoneSettingEditResponseResultZonesTLSClientAuthValueOff ZoneSettingEditResponseResultZonesTLSClientAuthValue = "off"
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
type ZoneSettingEditResponseResultZonesWAF struct {
	// ID of the zone setting.
	ID ZoneSettingEditResponseResultZonesWAFID `json:"id"`
	// Whether or not this setting can be modified for this zone (based on your
	// Cloudflare plan level).
	Editable ZoneSettingEditResponseResultZonesWAFEditable `json:"editable"`
	// last time this setting was modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Value of the zone setting.
	Value ZoneSettingEditResponseResultZonesWAFValue `json:"value"`
	JSON  zoneSettingEditResponseResultZonesWAFJSON  `json:"-"`
}

// zoneSettingEditResponseResultZonesWAFJSON contains the JSON metadata for the
// struct [ZoneSettingEditResponseResultZonesWAF]
type zoneSettingEditResponseResultZonesWAFJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSettingEditResponseResultZonesWAF) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneSettingEditResponseResultZonesWAF) implementsZoneSettingEditResponseResult() {}

// ID of the zone setting.
type ZoneSettingEditResponseResultZonesWAFID string

const (
	ZoneSettingEditResponseResultZonesWAFIDWAF ZoneSettingEditResponseResultZonesWAFID = "waf"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditResponseResultZonesWAFEditable bool

const (
	ZoneSettingEditResponseResultZonesWAFEditableTrue  ZoneSettingEditResponseResultZonesWAFEditable = true
	ZoneSettingEditResponseResultZonesWAFEditableFalse ZoneSettingEditResponseResultZonesWAFEditable = false
)

// Value of the zone setting.
type ZoneSettingEditResponseResultZonesWAFValue string

const (
	ZoneSettingEditResponseResultZonesWAFValueOn  ZoneSettingEditResponseResultZonesWAFValue = "on"
	ZoneSettingEditResponseResultZonesWAFValueOff ZoneSettingEditResponseResultZonesWAFValue = "off"
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
// [ZoneSettingEditParamsItemsZonesAdvancedDDOS],
// [ZoneSettingEditParamsItemsZonesAlwaysOnline],
// [ZoneSettingEditParamsItemsZonesAlwaysUseHTTPs],
// [ZoneSettingEditParamsItemsZonesAutomaticHTTPsRewrites],
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
type ZoneSettingEditParamsItemsZonesAdvancedDDOS struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesAdvancedDDOSID] `json:"id"`
	// Value of the zone setting. Notes: Defaults to on for Business+ plans
	Value param.Field[ZoneSettingEditParamsItemsZonesAdvancedDDOSValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesAdvancedDDOSEditable bool

const (
	ZoneSettingEditParamsItemsZonesAdvancedDDOSEditableTrue  ZoneSettingEditParamsItemsZonesAdvancedDDOSEditable = true
	ZoneSettingEditParamsItemsZonesAdvancedDDOSEditableFalse ZoneSettingEditParamsItemsZonesAdvancedDDOSEditable = false
)

// Value of the zone setting. Notes: Defaults to on for Business+ plans
type ZoneSettingEditParamsItemsZonesAdvancedDDOSValue string

const (
	ZoneSettingEditParamsItemsZonesAdvancedDDOSValueOn  ZoneSettingEditParamsItemsZonesAdvancedDDOSValue = "on"
	ZoneSettingEditParamsItemsZonesAdvancedDDOSValueOff ZoneSettingEditParamsItemsZonesAdvancedDDOSValue = "off"
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
type ZoneSettingEditParamsItemsZonesBrowserCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesBrowserCacheTTLID] `json:"id"`
	// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
	// `Respect Existing Headers`
	Value param.Field[ZoneSettingEditParamsItemsZonesBrowserCacheTTLValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditable bool

const (
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditableTrue  ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditable = true
	ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditableFalse ZoneSettingEditParamsItemsZonesBrowserCacheTTLEditable = false
)

// Value of the zone setting. Notes: Setting a TTL of 0 is equivalent to selecting
// `Respect Existing Headers`
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
type ZoneSettingEditParamsItemsZonesChallengeTTL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesChallengeTTLID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesChallengeTTLValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesChallengeTTLEditable bool

const (
	ZoneSettingEditParamsItemsZonesChallengeTTLEditableTrue  ZoneSettingEditParamsItemsZonesChallengeTTLEditable = true
	ZoneSettingEditParamsItemsZonesChallengeTTLEditableFalse ZoneSettingEditParamsItemsZonesChallengeTTLEditable = false
)

// Value of the zone setting.
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
type ZoneSettingEditParamsItemsZonesEdgeCacheTTL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesEdgeCacheTTLID] `json:"id"`
	// Value of the zone setting. Notes: The minimum TTL available depends on the plan
	// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
	Value param.Field[ZoneSettingEditParamsItemsZonesEdgeCacheTTLValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditable bool

const (
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditableTrue  ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditable = true
	ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditableFalse ZoneSettingEditParamsItemsZonesEdgeCacheTTLEditable = false
)

// Value of the zone setting. Notes: The minimum TTL available depends on the plan
// level of the zone. (Enterprise = 30, Business = 1800, Pro = 3600, Free = 7200)
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
type ZoneSettingEditParamsItemsZonesHTTP2 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHTTP2ID] `json:"id"`
	// Value of the HTTP2 setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHTTP2Value] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHTTP2Editable bool

const (
	ZoneSettingEditParamsItemsZonesHTTP2EditableTrue  ZoneSettingEditParamsItemsZonesHTTP2Editable = true
	ZoneSettingEditParamsItemsZonesHTTP2EditableFalse ZoneSettingEditParamsItemsZonesHTTP2Editable = false
)

// Value of the HTTP2 setting.
type ZoneSettingEditParamsItemsZonesHTTP2Value string

const (
	ZoneSettingEditParamsItemsZonesHTTP2ValueOn  ZoneSettingEditParamsItemsZonesHTTP2Value = "on"
	ZoneSettingEditParamsItemsZonesHTTP2ValueOff ZoneSettingEditParamsItemsZonesHTTP2Value = "off"
)

// HTTP3 enabled for this zone.
type ZoneSettingEditParamsItemsZonesHTTP3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesHTTP3ID] `json:"id"`
	// Value of the HTTP3 setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesHTTP3Value] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesHTTP3Editable bool

const (
	ZoneSettingEditParamsItemsZonesHTTP3EditableTrue  ZoneSettingEditParamsItemsZonesHTTP3Editable = true
	ZoneSettingEditParamsItemsZonesHTTP3EditableFalse ZoneSettingEditParamsItemsZonesHTTP3Editable = false
)

// Value of the HTTP3 setting.
type ZoneSettingEditParamsItemsZonesHTTP3Value string

const (
	ZoneSettingEditParamsItemsZonesHTTP3ValueOn  ZoneSettingEditParamsItemsZonesHTTP3Value = "on"
	ZoneSettingEditParamsItemsZonesHTTP3ValueOff ZoneSettingEditParamsItemsZonesHTTP3Value = "off"
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
type ZoneSettingEditParamsItemsZonesIPV6 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesIPV6ID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesIPV6Value] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesIPV6Editable bool

const (
	ZoneSettingEditParamsItemsZonesIPV6EditableTrue  ZoneSettingEditParamsItemsZonesIPV6Editable = true
	ZoneSettingEditParamsItemsZonesIPV6EditableFalse ZoneSettingEditParamsItemsZonesIPV6Editable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesIPV6Value string

const (
	ZoneSettingEditParamsItemsZonesIPV6ValueOff ZoneSettingEditParamsItemsZonesIPV6Value = "off"
	ZoneSettingEditParamsItemsZonesIPV6ValueOn  ZoneSettingEditParamsItemsZonesIPV6Value = "on"
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
type ZoneSettingEditParamsItemsZonesMinTLSVersion struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesMinTLSVersionID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesMinTLSVersionValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesMinTLSVersionEditable bool

const (
	ZoneSettingEditParamsItemsZonesMinTLSVersionEditableTrue  ZoneSettingEditParamsItemsZonesMinTLSVersionEditable = true
	ZoneSettingEditParamsItemsZonesMinTLSVersionEditableFalse ZoneSettingEditParamsItemsZonesMinTLSVersionEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesMinTLSVersionValue string

const (
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_0 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.0"
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_1 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.1"
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_2 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.2"
	ZoneSettingEditParamsItemsZonesMinTLSVersionValue1_3 ZoneSettingEditParamsItemsZonesMinTLSVersionValue = "1.3"
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
type ZoneSettingEditParamsItemsZonesNEL struct {
	// Zone setting identifier.
	ID param.Field[ZoneSettingEditParamsItemsZonesNELID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesNELValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesNELEditable bool

const (
	ZoneSettingEditParamsItemsZonesNELEditableTrue  ZoneSettingEditParamsItemsZonesNELEditable = true
	ZoneSettingEditParamsItemsZonesNELEditableFalse ZoneSettingEditParamsItemsZonesNELEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesNELValue struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneSettingEditParamsItemsZonesNELValue) MarshalJSON() (data []byte, err error) {
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
type ZoneSettingEditParamsItemsZonesSSL struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesSSLID] `json:"id"`
	// Value of the zone setting. Notes: Depends on the zone's plan level
	Value param.Field[ZoneSettingEditParamsItemsZonesSSLValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesSSLEditable bool

const (
	ZoneSettingEditParamsItemsZonesSSLEditableTrue  ZoneSettingEditParamsItemsZonesSSLEditable = true
	ZoneSettingEditParamsItemsZonesSSLEditableFalse ZoneSettingEditParamsItemsZonesSSLEditable = false
)

// Value of the zone setting. Notes: Depends on the zone's plan level
type ZoneSettingEditParamsItemsZonesSSLValue string

const (
	ZoneSettingEditParamsItemsZonesSSLValueOff      ZoneSettingEditParamsItemsZonesSSLValue = "off"
	ZoneSettingEditParamsItemsZonesSSLValueFlexible ZoneSettingEditParamsItemsZonesSSLValue = "flexible"
	ZoneSettingEditParamsItemsZonesSSLValueFull     ZoneSettingEditParamsItemsZonesSSLValue = "full"
	ZoneSettingEditParamsItemsZonesSSLValueStrict   ZoneSettingEditParamsItemsZonesSSLValue = "strict"
)

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
	ID param.Field[ZoneSettingEditParamsItemsZonesTLS1_2OnlyID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditable bool

const (
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditableTrue  ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditable = true
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditableFalse ZoneSettingEditParamsItemsZonesTLS1_2OnlyEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue string

const (
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyValueOff ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue = "off"
	ZoneSettingEditParamsItemsZonesTLS1_2OnlyValueOn  ZoneSettingEditParamsItemsZonesTLS1_2OnlyValue = "on"
)

// Enables Crypto TLS 1.3 feature for a zone.
type ZoneSettingEditParamsItemsZonesTLS1_3 struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTLS1_3ID] `json:"id"`
	// Value of the zone setting. Notes: Default value depends on the zone's plan
	// level.
	Value param.Field[ZoneSettingEditParamsItemsZonesTLS1_3Value] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTLS1_3Editable bool

const (
	ZoneSettingEditParamsItemsZonesTLS1_3EditableTrue  ZoneSettingEditParamsItemsZonesTLS1_3Editable = true
	ZoneSettingEditParamsItemsZonesTLS1_3EditableFalse ZoneSettingEditParamsItemsZonesTLS1_3Editable = false
)

// Value of the zone setting. Notes: Default value depends on the zone's plan
// level.
type ZoneSettingEditParamsItemsZonesTLS1_3Value string

const (
	ZoneSettingEditParamsItemsZonesTLS1_3ValueOn  ZoneSettingEditParamsItemsZonesTLS1_3Value = "on"
	ZoneSettingEditParamsItemsZonesTLS1_3ValueOff ZoneSettingEditParamsItemsZonesTLS1_3Value = "off"
	ZoneSettingEditParamsItemsZonesTLS1_3ValueZrt ZoneSettingEditParamsItemsZonesTLS1_3Value = "zrt"
)

// TLS Client Auth requires Cloudflare to connect to your origin server using a
// client certificate (Enterprise Only).
type ZoneSettingEditParamsItemsZonesTLSClientAuth struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesTLSClientAuthID] `json:"id"`
	// value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesTLSClientAuthValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesTLSClientAuthEditable bool

const (
	ZoneSettingEditParamsItemsZonesTLSClientAuthEditableTrue  ZoneSettingEditParamsItemsZonesTLSClientAuthEditable = true
	ZoneSettingEditParamsItemsZonesTLSClientAuthEditableFalse ZoneSettingEditParamsItemsZonesTLSClientAuthEditable = false
)

// value of the zone setting.
type ZoneSettingEditParamsItemsZonesTLSClientAuthValue string

const (
	ZoneSettingEditParamsItemsZonesTLSClientAuthValueOn  ZoneSettingEditParamsItemsZonesTLSClientAuthValue = "on"
	ZoneSettingEditParamsItemsZonesTLSClientAuthValueOff ZoneSettingEditParamsItemsZonesTLSClientAuthValue = "off"
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
type ZoneSettingEditParamsItemsZonesWAF struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsZonesWAFID] `json:"id"`
	// Value of the zone setting.
	Value param.Field[ZoneSettingEditParamsItemsZonesWAFValue] `json:"value"`
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

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsZonesWAFEditable bool

const (
	ZoneSettingEditParamsItemsZonesWAFEditableTrue  ZoneSettingEditParamsItemsZonesWAFEditable = true
	ZoneSettingEditParamsItemsZonesWAFEditableFalse ZoneSettingEditParamsItemsZonesWAFEditable = false
)

// Value of the zone setting.
type ZoneSettingEditParamsItemsZonesWAFValue string

const (
	ZoneSettingEditParamsItemsZonesWAFValueOn  ZoneSettingEditParamsItemsZonesWAFValue = "on"
	ZoneSettingEditParamsItemsZonesWAFValueOff ZoneSettingEditParamsItemsZonesWAFValue = "off"
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
