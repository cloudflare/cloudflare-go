package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
	"github.com/cloudflare/cloudflare-sdk-go/responses"
)

type ZonesSettingService struct {
	Options                        []options.RequestOption
	AdvancedDdos                   *ZonesSettingsAdvancedDdoService
	AlwaysOnlines                  *ZonesSettingsAlwaysOnlineService
	AlwaysUseHTTPs                 *ZonesSettingsAlwaysUseHTTPService
	AutomaticHTTPsRewrites         *ZonesSettingsAutomaticHTTPsRewriteService
	AutomaticPlatformOptimizations *ZonesSettingsAutomaticPlatformOptimizationService
	Brotlis                        *ZonesSettingsBrotliService
	BrowserCacheTtls               *ZonesSettingsBrowserCacheTtlService
	BrowserChecks                  *ZonesSettingsBrowserCheckService
	CacheLevels                    *ZonesSettingsCacheLevelService
	ChallengeTtls                  *ZonesSettingsChallengeTtlService
	Ciphers                        *ZonesSettingsCipherService
	DevelopmentModes               *ZonesSettingsDevelopmentModeService
	EarlyHints                     *ZonesSettingsEarlyHintService
	EmailObfuscations              *ZonesSettingsEmailObfuscationService
	H2Prioritizations              *ZonesSettingsH2PrioritizationService
	HotlinkProtections             *ZonesSettingsHotlinkProtectionService
	Http2s                         *ZonesSettingsHttp2Service
	Http3s                         *ZonesSettingsHttp3Service
	ImageResizings                 *ZonesSettingsImageResizingService
	IpGeolocations                 *ZonesSettingsIpGeolocationService
	Ipv6s                          *ZonesSettingsIpv6Service
	MinTlsVersions                 *ZonesSettingsMinTlsVersionService
	Minifies                       *ZonesSettingsMinifyService
	Mirages                        *ZonesSettingsMirageService
	MobileRedirects                *ZonesSettingsMobileRedirectService
	Nels                           *ZonesSettingsNelService
	OpportunisticEncryptions       *ZonesSettingsOpportunisticEncryptionService
	OpportunisticOnions            *ZonesSettingsOpportunisticOnionService
	OrangeToOranges                *ZonesSettingsOrangeToOrangeService
	OriginErrorPagePassThrus       *ZonesSettingsOriginErrorPagePassThrusService
	OriginMaxHTTPVersions          *ZonesSettingsOriginMaxHTTPVersionService
	Polishes                       *ZonesSettingsPolishService
	PrefetchPreloads               *ZonesSettingsPrefetchPreloadService
	PrivacyPasses                  *ZonesSettingsPrivacyPassService
	ProxyReadTimeouts              *ZonesSettingsProxyReadTimeoutService
	PseudoIpv4s                    *ZonesSettingsPseudoIpv4Service
	ResponseBufferings             *ZonesSettingsResponseBufferingService
	RocketLoaders                  *ZonesSettingsRocketLoaderService
	SecurityHeaders                *ZonesSettingsSecurityHeaderService
	SecurityLevels                 *ZonesSettingsSecurityLevelService
	ServerSideExcludes             *ZonesSettingsServerSideExcludeService
	SortQueryStringForCaches       *ZonesSettingsSortQueryStringForCachService
	Ssls                           *ZonesSettingsSslService
	SslRecommenders                *ZonesSettingsSslRecommenderService
	Tls_1_3s                       *ZonesSettingsTls_1_3Service
	TlsClientAuths                 *ZonesSettingsTlsClientAuthService
	TrueClientIpHeaders            *ZonesSettingsTrueClientIpHeaderService
	Wafs                           *ZonesSettingsWafService
	Webps                          *ZonesSettingsWebpService
	Websockets                     *ZonesSettingsWebsocketService
}

func NewZonesSettingService(opts ...options.RequestOption) (r *ZonesSettingService) {
	r = &ZonesSettingService{}
	r.Options = opts
	r.AdvancedDdos = NewZonesSettingsAdvancedDdoService(opts...)
	r.AlwaysOnlines = NewZonesSettingsAlwaysOnlineService(opts...)
	r.AlwaysUseHTTPs = NewZonesSettingsAlwaysUseHTTPService(opts...)
	r.AutomaticHTTPsRewrites = NewZonesSettingsAutomaticHTTPsRewriteService(opts...)
	r.AutomaticPlatformOptimizations = NewZonesSettingsAutomaticPlatformOptimizationService(opts...)
	r.Brotlis = NewZonesSettingsBrotliService(opts...)
	r.BrowserCacheTtls = NewZonesSettingsBrowserCacheTtlService(opts...)
	r.BrowserChecks = NewZonesSettingsBrowserCheckService(opts...)
	r.CacheLevels = NewZonesSettingsCacheLevelService(opts...)
	r.ChallengeTtls = NewZonesSettingsChallengeTtlService(opts...)
	r.Ciphers = NewZonesSettingsCipherService(opts...)
	r.DevelopmentModes = NewZonesSettingsDevelopmentModeService(opts...)
	r.EarlyHints = NewZonesSettingsEarlyHintService(opts...)
	r.EmailObfuscations = NewZonesSettingsEmailObfuscationService(opts...)
	r.H2Prioritizations = NewZonesSettingsH2PrioritizationService(opts...)
	r.HotlinkProtections = NewZonesSettingsHotlinkProtectionService(opts...)
	r.Http2s = NewZonesSettingsHttp2Service(opts...)
	r.Http3s = NewZonesSettingsHttp3Service(opts...)
	r.ImageResizings = NewZonesSettingsImageResizingService(opts...)
	r.IpGeolocations = NewZonesSettingsIpGeolocationService(opts...)
	r.Ipv6s = NewZonesSettingsIpv6Service(opts...)
	r.MinTlsVersions = NewZonesSettingsMinTlsVersionService(opts...)
	r.Minifies = NewZonesSettingsMinifyService(opts...)
	r.Mirages = NewZonesSettingsMirageService(opts...)
	r.MobileRedirects = NewZonesSettingsMobileRedirectService(opts...)
	r.Nels = NewZonesSettingsNelService(opts...)
	r.OpportunisticEncryptions = NewZonesSettingsOpportunisticEncryptionService(opts...)
	r.OpportunisticOnions = NewZonesSettingsOpportunisticOnionService(opts...)
	r.OrangeToOranges = NewZonesSettingsOrangeToOrangeService(opts...)
	r.OriginErrorPagePassThrus = NewZonesSettingsOriginErrorPagePassThrusService(opts...)
	r.OriginMaxHTTPVersions = NewZonesSettingsOriginMaxHTTPVersionService(opts...)
	r.Polishes = NewZonesSettingsPolishService(opts...)
	r.PrefetchPreloads = NewZonesSettingsPrefetchPreloadService(opts...)
	r.PrivacyPasses = NewZonesSettingsPrivacyPassService(opts...)
	r.ProxyReadTimeouts = NewZonesSettingsProxyReadTimeoutService(opts...)
	r.PseudoIpv4s = NewZonesSettingsPseudoIpv4Service(opts...)
	r.ResponseBufferings = NewZonesSettingsResponseBufferingService(opts...)
	r.RocketLoaders = NewZonesSettingsRocketLoaderService(opts...)
	r.SecurityHeaders = NewZonesSettingsSecurityHeaderService(opts...)
	r.SecurityLevels = NewZonesSettingsSecurityLevelService(opts...)
	r.ServerSideExcludes = NewZonesSettingsServerSideExcludeService(opts...)
	r.SortQueryStringForCaches = NewZonesSettingsSortQueryStringForCachService(opts...)
	r.Ssls = NewZonesSettingsSslService(opts...)
	r.SslRecommenders = NewZonesSettingsSslRecommenderService(opts...)
	r.Tls_1_3s = NewZonesSettingsTls_1_3Service(opts...)
	r.TlsClientAuths = NewZonesSettingsTlsClientAuthService(opts...)
	r.TrueClientIpHeaders = NewZonesSettingsTrueClientIpHeaderService(opts...)
	r.Wafs = NewZonesSettingsWafService(opts...)
	r.Webps = NewZonesSettingsWebpService(opts...)
	r.Websockets = NewZonesSettingsWebsocketService(opts...)
	return
}

// Available settings for your user in relation to a zone.
func (r *ZonesSettingService) List(ctx context.Context, zone_identifier string, opts ...options.RequestOption) (res *responses.ZoneSettingsCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "GET", path, nil, &res, opts...)
	return
}

// Edit settings for a zone.
func (r *ZonesSettingService) Edit(ctx context.Context, zone_identifier string, body *requests.SettingEditParams, opts ...options.RequestOption) (res *responses.ZoneSettingsCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/settings", zone_identifier)
	err = options.ExecuteNewRequest(ctx, "PATCH", path, body, &res, opts...)
	return
}
