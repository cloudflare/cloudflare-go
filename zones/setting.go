// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
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
