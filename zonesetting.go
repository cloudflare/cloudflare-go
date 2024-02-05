// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

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
	AdvancedDDOS                   *ZoneSettingAdvancedDDOSService
	AlwaysOnlines                  *ZoneSettingAlwaysOnlineService
	AlwaysUseHTTPs                 *ZoneSettingAlwaysUseHTTPService
	AutomaticHTTPsRewrites         *ZoneSettingAutomaticHTTPsRewriteService
	AutomaticPlatformOptimizations *ZoneSettingAutomaticPlatformOptimizationService
	Brotli                         *ZoneSettingBrotliService
	BrowserCacheTTLs               *ZoneSettingBrowserCacheTTLService
	BrowserChecks                  *ZoneSettingBrowserCheckService
	CacheLevels                    *ZoneSettingCacheLevelService
	ChallengeTTLs                  *ZoneSettingChallengeTTLService
	Ciphers                        *ZoneSettingCipherService
	DevelopmentModes               *ZoneSettingDevelopmentModeService
	EarlyHints                     *ZoneSettingEarlyHintService
	EmailObfuscations              *ZoneSettingEmailObfuscationService
	H2Prioritizations              *ZoneSettingH2PrioritizationService
	HotlinkProtections             *ZoneSettingHotlinkProtectionService
	HTTP2s                         *ZoneSettingHTTP2Service
	HTTP3s                         *ZoneSettingHTTP3Service
	ImageResizings                 *ZoneSettingImageResizingService
	IPGeolocations                 *ZoneSettingIPGeolocationService
	IPV6s                          *ZoneSettingIPV6Service
	MinTLSVersions                 *ZoneSettingMinTLSVersionService
	Minifies                       *ZoneSettingMinifyService
	Mirages                        *ZoneSettingMirageService
	MobileRedirects                *ZoneSettingMobileRedirectService
	NELs                           *ZoneSettingNELService
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
	SSLs                           *ZoneSettingSSLService
	SSLRecommenders                *ZoneSettingSSLRecommenderService
	TLS1_3s                        *ZoneSettingTLS1_3Service
	TLSClientAuths                 *ZoneSettingTLSClientAuthService
	TrueClientIPHeaders            *ZoneSettingTrueClientIPHeaderService
	WAFs                           *ZoneSettingWAFService
	Webps                          *ZoneSettingWebpService
	Websockets                     *ZoneSettingWebsocketService
}

// NewZoneSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSettingService(opts ...option.RequestOption) (r *ZoneSettingService) {
	r = &ZoneSettingService{}
	r.Options = opts
	r.AdvancedDDOS = NewZoneSettingAdvancedDDOSService(opts...)
	r.AlwaysOnlines = NewZoneSettingAlwaysOnlineService(opts...)
	r.AlwaysUseHTTPs = NewZoneSettingAlwaysUseHTTPService(opts...)
	r.AutomaticHTTPsRewrites = NewZoneSettingAutomaticHTTPsRewriteService(opts...)
	r.AutomaticPlatformOptimizations = NewZoneSettingAutomaticPlatformOptimizationService(opts...)
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
	Errors   []ZoneSettingListResponseError   `json:"errors,required"`
	Messages []ZoneSettingListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                        `json:"success,required"`
	JSON    zoneSettingListResponseJSON `json:"-"`
}

// zoneSettingListResponseJSON contains the JSON metadata for the struct
// [ZoneSettingListResponse]
type zoneSettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
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

type ZoneSettingEditResponse struct {
	Errors   []ZoneSettingEditResponseError   `json:"errors,required"`
	Messages []ZoneSettingEditResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success bool                        `json:"success,required"`
	JSON    zoneSettingEditResponseJSON `json:"-"`
}

// zoneSettingEditResponseJSON contains the JSON metadata for the struct
// [ZoneSettingEditResponse]
type zoneSettingEditResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
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

type ZoneSettingEditParams struct {
	// One or more zone setting objects. Must contain an ID and a value.
	Items param.Field[[]ZoneSettingEditParamsItem] `json:"items,required"`
}

func (r ZoneSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The highest HTTP version Cloudflare will attempt to use with your origin. This
// setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.).
//
// Satisfied by [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsZonesOriginMaxHTTPVersion],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsZonesSSLRecommender],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject], [ZoneSettingEditParamsItemsObject],
// [ZoneSettingEditParamsItemsObject].
type ZoneSettingEditParamsItem interface {
	implementsZoneSettingEditParamsItem()
}

type ZoneSettingEditParamsItemsObject struct {
	// ID of the zone setting.
	ID param.Field[ZoneSettingEditParamsItemsObjectID] `json:"id,required"`
	// Value of the 0-RTT setting.
	Value param.Field[ZoneSettingEditParamsItemsObjectValue] `json:"value,required"`
}

func (r ZoneSettingEditParamsItemsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneSettingEditParamsItemsObject) implementsZoneSettingEditParamsItem() {}

// ID of the zone setting.
type ZoneSettingEditParamsItemsObjectID string

const (
	ZoneSettingEditParamsItemsObjectID0rtt ZoneSettingEditParamsItemsObjectID = "0rtt"
)

// Value of the 0-RTT setting.
type ZoneSettingEditParamsItemsObjectValue string

const (
	ZoneSettingEditParamsItemsObjectValueOn  ZoneSettingEditParamsItemsObjectValue = "on"
	ZoneSettingEditParamsItemsObjectValueOff ZoneSettingEditParamsItemsObjectValue = "off"
)

// Whether or not this setting can be modified for this zone (based on your
// Cloudflare plan level).
type ZoneSettingEditParamsItemsObjectEditable bool

const (
	ZoneSettingEditParamsItemsObjectEditableTrue  ZoneSettingEditParamsItemsObjectEditable = true
	ZoneSettingEditParamsItemsObjectEditableFalse ZoneSettingEditParamsItemsObjectEditable = false
)

// The highest HTTP version Cloudflare will attempt to use with your origin. This
// setting allows Cloudflare to make HTTP/2 requests to your origin. (Refer to
// [Enable HTTP/2 to Origin](https://developers.cloudflare.com/cache/how-to/enable-http2-to-origin/),
// for more information.).
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
