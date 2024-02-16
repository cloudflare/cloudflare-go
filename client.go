// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"os"

	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                      []option.RequestOption
	Accounts                     *AccountService
	Certificates                 *CertificateService
	IPs                          *IPService
	Memberships                  *MembershipService
	Users                        *UserService
	Zones                        *ZoneService
	AI                           *AIService
	LoadBalancers                *LoadBalancerService
	Access                       *AccessService
	DNSAnalytics                 *DNSAnalyticService
	PurgeCaches                  *PurgeCachService
	SSLs                         *SSLService
	Subscriptions                *SubscriptionService
	Acms                         *AcmService
	Analytics                    *AnalyticsService
	Argo                         *ArgoService
	AvailablePlans               *AvailablePlanService
	AvailableRatePlans           *AvailableRatePlanService
	Caches                       *CachService
	CertificateAuthorities       *CertificateAuthorityService
	ClientCertificates           *ClientCertificateService
	CustomCertificates           *CustomCertificateService
	CustomHostnames              *CustomHostnameService
	CustomNs                     *CustomNService
	DNSRecords                   *DNSRecordService
	DNSSECs                      *DNSSECService
	Emails                       *EmailService
	Filters                      *FilterService
	Firewalls                    *FirewallService
	Healthchecks                 *HealthcheckService
	KeylessCertificates          *KeylessCertificateService
	Logpush                      *LogpushService
	Logs                         *LogService
	OriginTLSClientAuth          *OriginTLSClientAuthService
	Pagerules                    *PageruleService
	RateLimits                   *RateLimitService
	SecondaryDNS                 *SecondaryDNSService
	Settings                     *SettingService
	WaitingRooms                 *WaitingRoomService
	Web3s                        *Web3Service
	Workers                      *WorkerService
	ActivationChecks             *ActivationCheckService
	ManagedHeaders               *ManagedHeaderService
	PageShields                  *PageShieldService
	Rulesets                     *RulesetService
	URLNormalizations            *URLNormalizationService
	Spectrums                    *SpectrumService
	Addresses                    *AddressService
	AuditLogs                    *AuditLogService
	Billings                     *BillingService
	BrandProtections             *BrandProtectionService
	CfdTunnels                   *CfdTunnelService
	Diagnostics                  *DiagnosticService
	DLPs                         *DLPService
	DNSFirewalls                 *DNSFirewallService
	Images                       *ImageService
	Intels                       *IntelService
	Magics                       *MagicService
	AccountMembers               *AccountMemberService
	Mnms                         *MnmService
	MtlsCertificates             *MtlsCertificateService
	Pages                        *PageService
	Pcaps                        *PcapService
	Registrar                    *RegistrarService
	RequestTracers               *RequestTracerService
	Roles                        *RoleService
	Rules                        *RuleService
	Storage                      *StorageService
	Stream                       *StreamService
	Teamnets                     *TeamnetService
	Tunnels                      *TunnelService
	Gateways                     *GatewayService
	Alerting                     *AlertingService
	Devices                      *DeviceService
	D1                           *D1Service
	DEX                          *DEXService
	R2                           *R2Service
	Teamnet                      *TeamnetService
	WarpConnector                *WarpConnectorService
	Dispatchers                  *DispatcherService
	WorkersForPlatforms          *WorkersForPlatformService
	WorkerDomains                *WorkerDomainService
	WorkerScripts                *WorkerScriptService
	Zerotrust                    *ZerotrustService
	Addressing                   *AddressingService
	Challenges                   *ChallengeService
	Hyperdrive                   *HyperdriveService
	Intel                        *IntelService
	Rum                          *RumService
	Vectorize                    *VectorizeService
	URLScanner                   *URLScannerService
	Radar                        *RadarService
	BotManagements               *BotManagementService
	CacheReserves                *CacheReserveService
	OriginPostQuantumEncryptions *OriginPostQuantumEncryptionService
	Cache                        *CacheService
	Firewall                     *FirewallService
	Zaraz                        *ZarazService
	SpeedAPI                     *SpeedAPIService
	DcvDelegation                *DcvDelegationService
	Hostnames                    *HostnameService
	PageShield                   *PageShieldService
	FontSettings                 *FontSettingService
	Snippets                     *SnippetService
	DLP                          *DLPService
	Gateway                      *GatewayService
	AccessTags                   *AccessTagService
	Calls                        *CallService
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_KEY, CLOUDFLARE_EMAIL, CLOUDFLARE_API_TOKEN,
// CLOUDFLARE_API_USER_SERVICE_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_EMAIL"); ok {
		defaults = append(defaults, option.WithAPIEmail(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_TOKEN"); ok {
		defaults = append(defaults, option.WithAPIToken(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_USER_SERVICE_KEY"); ok {
		defaults = append(defaults, option.WithUserServiceKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.Certificates = NewCertificateService(opts...)
	r.IPs = NewIPService(opts...)
	r.Memberships = NewMembershipService(opts...)
	r.Users = NewUserService(opts...)
	r.Zones = NewZoneService(opts...)
	r.AI = NewAIService(opts...)
	r.LoadBalancers = NewLoadBalancerService(opts...)
	r.Access = NewAccessService(opts...)
	r.DNSAnalytics = NewDNSAnalyticService(opts...)
	r.PurgeCaches = NewPurgeCachService(opts...)
	r.SSLs = NewSSLService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Acms = NewAcmService(opts...)
	r.Analytics = NewAnalyticsService(opts...)
	r.Argo = NewArgoService(opts...)
	r.AvailablePlans = NewAvailablePlanService(opts...)
	r.AvailableRatePlans = NewAvailableRatePlanService(opts...)
	r.Caches = NewCachService(opts...)
	r.CertificateAuthorities = NewCertificateAuthorityService(opts...)
	r.ClientCertificates = NewClientCertificateService(opts...)
	r.CustomCertificates = NewCustomCertificateService(opts...)
	r.CustomHostnames = NewCustomHostnameService(opts...)
	r.CustomNs = NewCustomNService(opts...)
	r.DNSRecords = NewDNSRecordService(opts...)
	r.DNSSECs = NewDNSSECService(opts...)
	r.Emails = NewEmailService(opts...)
	r.Filters = NewFilterService(opts...)
	r.Firewalls = NewFirewallService(opts...)
	r.Healthchecks = NewHealthcheckService(opts...)
	r.KeylessCertificates = NewKeylessCertificateService(opts...)
	r.Logpush = NewLogpushService(opts...)
	r.Logs = NewLogService(opts...)
	r.OriginTLSClientAuth = NewOriginTLSClientAuthService(opts...)
	r.Pagerules = NewPageruleService(opts...)
	r.RateLimits = NewRateLimitService(opts...)
	r.SecondaryDNS = NewSecondaryDNSService(opts...)
	r.Settings = NewSettingService(opts...)
	r.WaitingRooms = NewWaitingRoomService(opts...)
	r.Web3s = NewWeb3Service(opts...)
	r.Workers = NewWorkerService(opts...)
	r.ActivationChecks = NewActivationCheckService(opts...)
	r.ManagedHeaders = NewManagedHeaderService(opts...)
	r.PageShields = NewPageShieldService(opts...)
	r.Rulesets = NewRulesetService(opts...)
	r.URLNormalizations = NewURLNormalizationService(opts...)
	r.Spectrums = NewSpectrumService(opts...)
	r.Addresses = NewAddressService(opts...)
	r.AuditLogs = NewAuditLogService(opts...)
	r.Billings = NewBillingService(opts...)
	r.BrandProtections = NewBrandProtectionService(opts...)
	r.CfdTunnels = NewCfdTunnelService(opts...)
	r.Diagnostics = NewDiagnosticService(opts...)
	r.DLPs = NewDLPService(opts...)
	r.DNSFirewalls = NewDNSFirewallService(opts...)
	r.Images = NewImageService(opts...)
	r.Intels = NewIntelService(opts...)
	r.Magics = NewMagicService(opts...)
	r.AccountMembers = NewAccountMemberService(opts...)
	r.Mnms = NewMnmService(opts...)
	r.MtlsCertificates = NewMtlsCertificateService(opts...)
	r.Pages = NewPageService(opts...)
	r.Pcaps = NewPcapService(opts...)
	r.Registrar = NewRegistrarService(opts...)
	r.RequestTracers = NewRequestTracerService(opts...)
	r.Roles = NewRoleService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Storage = NewStorageService(opts...)
	r.Stream = NewStreamService(opts...)
	r.Teamnets = NewTeamnetService(opts...)
	r.Tunnels = NewTunnelService(opts...)
	r.Gateways = NewGatewayService(opts...)
	r.Alerting = NewAlertingService(opts...)
	r.Devices = NewDeviceService(opts...)
	r.D1 = NewD1Service(opts...)
	r.DEX = NewDEXService(opts...)
	r.R2 = NewR2Service(opts...)
	r.Teamnet = NewTeamnetService(opts...)
	r.WarpConnector = NewWarpConnectorService(opts...)
	r.Dispatchers = NewDispatcherService(opts...)
	r.WorkersForPlatforms = NewWorkersForPlatformService(opts...)
	r.WorkerDomains = NewWorkerDomainService(opts...)
	r.WorkerScripts = NewWorkerScriptService(opts...)
	r.Zerotrust = NewZerotrustService(opts...)
	r.Addressing = NewAddressingService(opts...)
	r.Challenges = NewChallengeService(opts...)
	r.Hyperdrive = NewHyperdriveService(opts...)
	r.Intel = NewIntelService(opts...)
	r.Rum = NewRumService(opts...)
	r.Vectorize = NewVectorizeService(opts...)
	r.URLScanner = NewURLScannerService(opts...)
	r.Radar = NewRadarService(opts...)
	r.BotManagements = NewBotManagementService(opts...)
	r.CacheReserves = NewCacheReserveService(opts...)
	r.OriginPostQuantumEncryptions = NewOriginPostQuantumEncryptionService(opts...)
	r.Cache = NewCacheService(opts...)
	r.Firewall = NewFirewallService(opts...)
	r.Zaraz = NewZarazService(opts...)
	r.SpeedAPI = NewSpeedAPIService(opts...)
	r.DcvDelegation = NewDcvDelegationService(opts...)
	r.Hostnames = NewHostnameService(opts...)
	r.PageShield = NewPageShieldService(opts...)
	r.FontSettings = NewFontSettingService(opts...)
	r.Snippets = NewSnippetService(opts...)
	r.DLP = NewDLPService(opts...)
	r.Gateway = NewGatewayService(opts...)
	r.AccessTags = NewAccessTagService(opts...)
	r.Calls = NewCallService(opts...)

	return
}
