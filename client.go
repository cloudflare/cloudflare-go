// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"os"

	"github.com/cloudflare/cloudflare-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                     []option.RequestOption
	Accounts                    *AccountService
	Certificates                *CertificateService
	IPs                         *IPService
	Memberships                 *MembershipService
	User                        *UserService
	Zones                       *ZoneService
	LoadBalancers               *LoadBalancerService
	Cache                       *CacheService
	SSL                         *SSLService
	Subscriptions               *SubscriptionService
	ACM                         *ACMService
	Argo                        *ArgoService
	AvailablePlans              *AvailablePlanService
	AvailableRatePlans          *AvailableRatePlanService
	CertificateAuthorities      *CertificateAuthorityService
	ClientCertificates          *ClientCertificateService
	CustomCertificates          *CustomCertificateService
	CustomHostnames             *CustomHostnameService
	CustomNameservers           *CustomNameserverService
	DNS                         *DNSService
	DNSSEC                      *DNSSECService
	EmailRouting                *EmailRoutingService
	Filters                     *FilterService
	Firewall                    *FirewallService
	Healthchecks                *HealthcheckService
	KeylessCertificates         *KeylessCertificateService
	Logpush                     *LogpushService
	Logs                        *LogService
	OriginTLSClientAuth         *OriginTLSClientAuthService
	Pagerules                   *PageruleService
	RateLimits                  *RateLimitService
	SecondaryDNS                *SecondaryDNSService
	WaitingRooms                *WaitingRoomService
	Web3                        *Web3Service
	Workers                     *WorkerService
	KV                          *KVService
	DurableObjects              *DurableObjectService
	Queues                      *QueueService
	ManagedHeaders              *ManagedHeaderService
	PageShield                  *PageShieldService
	Rulesets                    *RulesetService
	URLNormalization            *URLNormalizationService
	Spectrum                    *SpectrumService
	Addressing                  *AddressingService
	AuditLogs                   *AuditLogService
	Billing                     *BillingService
	BrandProtection             *BrandProtectionService
	Diagnostics                 *DiagnosticService
	Images                      *ImageService
	Intel                       *IntelService
	MagicTransit                *MagicTransitService
	MagicNetworkMonitoring      *MagicNetworkMonitoringService
	MTLSCertificates            *MTLSCertificateService
	Pages                       *PageService
	PCAPs                       *PCAPService
	Registrar                   *RegistrarService
	RequestTracers              *RequestTracerService
	Rules                       *RuleService
	Storage                     *StorageService
	Stream                      *StreamService
	Alerting                    *AlertingService
	D1                          *D1Service
	R2                          *R2Service
	WARPConnector               *WARPConnectorService
	WorkersForPlatforms         *WorkersForPlatformService
	ZeroTrust                   *ZeroTrustService
	Challenges                  *ChallengeService
	Hyperdrive                  *HyperdriveService
	RUM                         *RUMService
	Vectorize                   *VectorizeService
	URLScanner                  *URLScannerService
	Radar                       *RadarService
	BotManagement               *BotManagementService
	OriginPostQuantumEncryption *OriginPostQuantumEncryptionService
	Speed                       *SpeedService
	DCVDelegation               *DCVDelegationService
	Hostnames                   *HostnameService
	Snippets                    *SnippetService
	Calls                       *CallService
	CloudforceOne               *CloudforceOneService
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
	r.User = NewUserService(opts...)
	r.Zones = NewZoneService(opts...)
	r.LoadBalancers = NewLoadBalancerService(opts...)
	r.Cache = NewCacheService(opts...)
	r.SSL = NewSSLService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.ACM = NewACMService(opts...)
	r.Argo = NewArgoService(opts...)
	r.AvailablePlans = NewAvailablePlanService(opts...)
	r.AvailableRatePlans = NewAvailableRatePlanService(opts...)
	r.CertificateAuthorities = NewCertificateAuthorityService(opts...)
	r.ClientCertificates = NewClientCertificateService(opts...)
	r.CustomCertificates = NewCustomCertificateService(opts...)
	r.CustomHostnames = NewCustomHostnameService(opts...)
	r.CustomNameservers = NewCustomNameserverService(opts...)
	r.DNS = NewDNSService(opts...)
	r.DNSSEC = NewDNSSECService(opts...)
	r.EmailRouting = NewEmailRoutingService(opts...)
	r.Filters = NewFilterService(opts...)
	r.Firewall = NewFirewallService(opts...)
	r.Healthchecks = NewHealthcheckService(opts...)
	r.KeylessCertificates = NewKeylessCertificateService(opts...)
	r.Logpush = NewLogpushService(opts...)
	r.Logs = NewLogService(opts...)
	r.OriginTLSClientAuth = NewOriginTLSClientAuthService(opts...)
	r.Pagerules = NewPageruleService(opts...)
	r.RateLimits = NewRateLimitService(opts...)
	r.SecondaryDNS = NewSecondaryDNSService(opts...)
	r.WaitingRooms = NewWaitingRoomService(opts...)
	r.Web3 = NewWeb3Service(opts...)
	r.Workers = NewWorkerService(opts...)
	r.KV = NewKVService(opts...)
	r.DurableObjects = NewDurableObjectService(opts...)
	r.Queues = NewQueueService(opts...)
	r.ManagedHeaders = NewManagedHeaderService(opts...)
	r.PageShield = NewPageShieldService(opts...)
	r.Rulesets = NewRulesetService(opts...)
	r.URLNormalization = NewURLNormalizationService(opts...)
	r.Spectrum = NewSpectrumService(opts...)
	r.Addressing = NewAddressingService(opts...)
	r.AuditLogs = NewAuditLogService(opts...)
	r.Billing = NewBillingService(opts...)
	r.BrandProtection = NewBrandProtectionService(opts...)
	r.Diagnostics = NewDiagnosticService(opts...)
	r.Images = NewImageService(opts...)
	r.Intel = NewIntelService(opts...)
	r.MagicTransit = NewMagicTransitService(opts...)
	r.MagicNetworkMonitoring = NewMagicNetworkMonitoringService(opts...)
	r.MTLSCertificates = NewMTLSCertificateService(opts...)
	r.Pages = NewPageService(opts...)
	r.PCAPs = NewPCAPService(opts...)
	r.Registrar = NewRegistrarService(opts...)
	r.RequestTracers = NewRequestTracerService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Storage = NewStorageService(opts...)
	r.Stream = NewStreamService(opts...)
	r.Alerting = NewAlertingService(opts...)
	r.D1 = NewD1Service(opts...)
	r.R2 = NewR2Service(opts...)
	r.WARPConnector = NewWARPConnectorService(opts...)
	r.WorkersForPlatforms = NewWorkersForPlatformService(opts...)
	r.ZeroTrust = NewZeroTrustService(opts...)
	r.Challenges = NewChallengeService(opts...)
	r.Hyperdrive = NewHyperdriveService(opts...)
	r.RUM = NewRUMService(opts...)
	r.Vectorize = NewVectorizeService(opts...)
	r.URLScanner = NewURLScannerService(opts...)
	r.Radar = NewRadarService(opts...)
	r.BotManagement = NewBotManagementService(opts...)
	r.OriginPostQuantumEncryption = NewOriginPostQuantumEncryptionService(opts...)
	r.Speed = NewSpeedService(opts...)
	r.DCVDelegation = NewDCVDelegationService(opts...)
	r.Hostnames = NewHostnameService(opts...)
	r.Snippets = NewSnippetService(opts...)
	r.Calls = NewCallService(opts...)
	r.CloudforceOne = NewCloudforceOneService(opts...)

	return
}
