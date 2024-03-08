// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"os"

	"github.com/cloudflare/cloudflare-go/accounts"
	"github.com/cloudflare/cloudflare-go/acm"
	"github.com/cloudflare/cloudflare-go/addressing"
	"github.com/cloudflare/cloudflare-go/alerting"
	"github.com/cloudflare/cloudflare-go/argo"
	"github.com/cloudflare/cloudflare-go/audit_logs"
	"github.com/cloudflare/cloudflare-go/available_plans"
	"github.com/cloudflare/cloudflare-go/available_rate_plans"
	"github.com/cloudflare/cloudflare-go/billing"
	"github.com/cloudflare/cloudflare-go/bot_management"
	"github.com/cloudflare/cloudflare-go/brand_protection"
	"github.com/cloudflare/cloudflare-go/cache"
	"github.com/cloudflare/cloudflare-go/calls"
	"github.com/cloudflare/cloudflare-go/certificate_authorities"
	"github.com/cloudflare/cloudflare-go/certificates"
	"github.com/cloudflare/cloudflare-go/challenges"
	"github.com/cloudflare/cloudflare-go/client_certificates"
	"github.com/cloudflare/cloudflare-go/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/custom_certificates"
	"github.com/cloudflare/cloudflare-go/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/custom_nameservers"
	"github.com/cloudflare/cloudflare-go/d1"
	"github.com/cloudflare/cloudflare-go/dcv_delegation"
	"github.com/cloudflare/cloudflare-go/diagnostics"
	"github.com/cloudflare/cloudflare-go/dns"
	"github.com/cloudflare/cloudflare-go/dnssec"
	"github.com/cloudflare/cloudflare-go/durable_objects"
	"github.com/cloudflare/cloudflare-go/email_routing"
	"github.com/cloudflare/cloudflare-go/filters"
	"github.com/cloudflare/cloudflare-go/firewall"
	"github.com/cloudflare/cloudflare-go/healthchecks"
	"github.com/cloudflare/cloudflare-go/hostnames"
	"github.com/cloudflare/cloudflare-go/hyperdrive"
	"github.com/cloudflare/cloudflare-go/images"
	"github.com/cloudflare/cloudflare-go/intel"
	"github.com/cloudflare/cloudflare-go/ips"
	"github.com/cloudflare/cloudflare-go/keyless_certificates"
	"github.com/cloudflare/cloudflare-go/kv"
	"github.com/cloudflare/cloudflare-go/load_balancers"
	"github.com/cloudflare/cloudflare-go/logpush"
	"github.com/cloudflare/cloudflare-go/logs"
	"github.com/cloudflare/cloudflare-go/magic_network_monitoring"
	"github.com/cloudflare/cloudflare-go/magic_transit"
	"github.com/cloudflare/cloudflare-go/managed_headers"
	"github.com/cloudflare/cloudflare-go/memberships"
	"github.com/cloudflare/cloudflare-go/mtls_certificates"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/origin_post_quantum_encryption"
	"github.com/cloudflare/cloudflare-go/origin_tls_client_auth"
	"github.com/cloudflare/cloudflare-go/page_shield"
	"github.com/cloudflare/cloudflare-go/pagerules"
	"github.com/cloudflare/cloudflare-go/pages"
	"github.com/cloudflare/cloudflare-go/pcaps"
	"github.com/cloudflare/cloudflare-go/queues"
	"github.com/cloudflare/cloudflare-go/r2"
	"github.com/cloudflare/cloudflare-go/radar"
	"github.com/cloudflare/cloudflare-go/rate_limits"
	"github.com/cloudflare/cloudflare-go/registrar"
	"github.com/cloudflare/cloudflare-go/request_tracers"
	"github.com/cloudflare/cloudflare-go/rules"
	"github.com/cloudflare/cloudflare-go/rulesets"
	"github.com/cloudflare/cloudflare-go/rum"
	"github.com/cloudflare/cloudflare-go/secondary_dns"
	"github.com/cloudflare/cloudflare-go/snippets"
	"github.com/cloudflare/cloudflare-go/spectrum"
	"github.com/cloudflare/cloudflare-go/speed"
	"github.com/cloudflare/cloudflare-go/ssl"
	"github.com/cloudflare/cloudflare-go/storage"
	"github.com/cloudflare/cloudflare-go/stream"
	"github.com/cloudflare/cloudflare-go/subscriptions"
	"github.com/cloudflare/cloudflare-go/url_normalization"
	"github.com/cloudflare/cloudflare-go/url_scanner"
	"github.com/cloudflare/cloudflare-go/user"
	"github.com/cloudflare/cloudflare-go/vectorize"
	"github.com/cloudflare/cloudflare-go/waiting_rooms"
	"github.com/cloudflare/cloudflare-go/warp_connector"
	"github.com/cloudflare/cloudflare-go/web3"
	"github.com/cloudflare/cloudflare-go/workers"
	"github.com/cloudflare/cloudflare-go/workers_for_platforms"
	"github.com/cloudflare/cloudflare-go/zero_trust"
	"github.com/cloudflare/cloudflare-go/zones"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                     []option.RequestOption
	Accounts                    *accounts.AccountService
	Certificates                *certificates.CertificateService
	IPs                         *ips.IPService
	Memberships                 *memberships.MembershipService
	User                        *user.UserService
	Zones                       *zones.ZoneService
	LoadBalancers               *load_balancers.LoadBalancerService
	Cache                       *cache.CacheService
	SSL                         *ssl.SSLService
	Subscriptions               *subscriptions.SubscriptionService
	ACM                         *acm.ACMService
	Argo                        *argo.ArgoService
	AvailablePlans              *available_plans.AvailablePlanService
	AvailableRatePlans          *available_rate_plans.AvailableRatePlanService
	CertificateAuthorities      *certificate_authorities.CertificateAuthorityService
	ClientCertificates          *client_certificates.ClientCertificateService
	CustomCertificates          *custom_certificates.CustomCertificateService
	CustomHostnames             *custom_hostnames.CustomHostnameService
	CustomNameservers           *custom_nameservers.CustomNameserverService
	DNS                         *dns.DNSService
	DNSSEC                      *dnssec.DNSSECService
	EmailRouting                *email_routing.EmailRoutingService
	Filters                     *filters.FilterService
	Firewall                    *firewall.FirewallService
	Healthchecks                *healthchecks.HealthcheckService
	KeylessCertificates         *keyless_certificates.KeylessCertificateService
	Logpush                     *logpush.LogpushService
	Logs                        *logs.LogService
	OriginTLSClientAuth         *origin_tls_client_auth.OriginTLSClientAuthService
	Pagerules                   *pagerules.PageruleService
	RateLimits                  *rate_limits.RateLimitService
	SecondaryDNS                *secondary_dns.SecondaryDNSService
	WaitingRooms                *waiting_rooms.WaitingRoomService
	Web3                        *web3.Web3Service
	Workers                     *workers.WorkerService
	KV                          *kv.KVService
	DurableObjects              *durable_objects.DurableObjectService
	Queues                      *queues.QueueService
	ManagedHeaders              *managed_headers.ManagedHeaderService
	PageShield                  *page_shield.PageShieldService
	Rulesets                    *rulesets.RulesetService
	URLNormalization            *url_normalization.URLNormalizationService
	Spectrum                    *spectrum.SpectrumService
	Addressing                  *addressing.AddressingService
	AuditLogs                   *audit_logs.AuditLogService
	Billing                     *billing.BillingService
	BrandProtection             *brand_protection.BrandProtectionService
	Diagnostics                 *diagnostics.DiagnosticService
	Images                      *images.ImageService
	Intel                       *intel.IntelService
	MagicTransit                *magic_transit.MagicTransitService
	MagicNetworkMonitoring      *magic_network_monitoring.MagicNetworkMonitoringService
	MTLSCertificates            *mtls_certificates.MTLSCertificateService
	Pages                       *pages.PageService
	PCAPs                       *pcaps.PCAPService
	Registrar                   *registrar.RegistrarService
	RequestTracers              *request_tracers.RequestTracerService
	Rules                       *rules.RuleService
	Storage                     *storage.StorageService
	Stream                      *stream.StreamService
	Alerting                    *alerting.AlertingService
	D1                          *d1.D1Service
	R2                          *r2.R2Service
	WARPConnector               *warp_connector.WARPConnectorService
	WorkersForPlatforms         *workers_for_platforms.WorkersForPlatformService
	ZeroTrust                   *zero_trust.ZeroTrustService
	Challenges                  *challenges.ChallengeService
	Hyperdrive                  *hyperdrive.HyperdriveService
	RUM                         *rum.RUMService
	Vectorize                   *vectorize.VectorizeService
	URLScanner                  *url_scanner.URLScannerService
	Radar                       *radar.RadarService
	BotManagement               *bot_management.BotManagementService
	OriginPostQuantumEncryption *origin_post_quantum_encryption.OriginPostQuantumEncryptionService
	Speed                       *speed.SpeedService
	DCVDelegation               *dcv_delegation.DCVDelegationService
	Hostnames                   *hostnames.HostnameService
	Snippets                    *snippets.SnippetService
	Calls                       *calls.CallService
	CloudforceOne               *cloudforce_one.CloudforceOneService
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

	r.Accounts = accounts.NewAccountService(opts...)
	r.Certificates = certificates.NewCertificateService(opts...)
	r.IPs = ips.NewIPService(opts...)
	r.Memberships = memberships.NewMembershipService(opts...)
	r.User = user.NewUserService(opts...)
	r.Zones = zones.NewZoneService(opts...)
	r.LoadBalancers = load_balancers.NewLoadBalancerService(opts...)
	r.Cache = cache.NewCacheService(opts...)
	r.SSL = ssl.NewSSLService(opts...)
	r.Subscriptions = subscriptions.NewSubscriptionService(opts...)
	r.ACM = acm.NewACMService(opts...)
	r.Argo = argo.NewArgoService(opts...)
	r.AvailablePlans = available_plans.NewAvailablePlanService(opts...)
	r.AvailableRatePlans = available_rate_plans.NewAvailableRatePlanService(opts...)
	r.CertificateAuthorities = certificate_authorities.NewCertificateAuthorityService(opts...)
	r.ClientCertificates = client_certificates.NewClientCertificateService(opts...)
	r.CustomCertificates = custom_certificates.NewCustomCertificateService(opts...)
	r.CustomHostnames = custom_hostnames.NewCustomHostnameService(opts...)
	r.CustomNameservers = custom_nameservers.NewCustomNameserverService(opts...)
	r.DNS = dns.NewDNSService(opts...)
	r.DNSSEC = dnssec.NewDNSSECService(opts...)
	r.EmailRouting = email_routing.NewEmailRoutingService(opts...)
	r.Filters = filters.NewFilterService(opts...)
	r.Firewall = firewall.NewFirewallService(opts...)
	r.Healthchecks = healthchecks.NewHealthcheckService(opts...)
	r.KeylessCertificates = keyless_certificates.NewKeylessCertificateService(opts...)
	r.Logpush = logpush.NewLogpushService(opts...)
	r.Logs = logs.NewLogService(opts...)
	r.OriginTLSClientAuth = origin_tls_client_auth.NewOriginTLSClientAuthService(opts...)
	r.Pagerules = pagerules.NewPageruleService(opts...)
	r.RateLimits = rate_limits.NewRateLimitService(opts...)
	r.SecondaryDNS = secondary_dns.NewSecondaryDNSService(opts...)
	r.WaitingRooms = waiting_rooms.NewWaitingRoomService(opts...)
	r.Web3 = web3.NewWeb3Service(opts...)
	r.Workers = workers.NewWorkerService(opts...)
	r.KV = kv.NewKVService(opts...)
	r.DurableObjects = durable_objects.NewDurableObjectService(opts...)
	r.Queues = queues.NewQueueService(opts...)
	r.ManagedHeaders = managed_headers.NewManagedHeaderService(opts...)
	r.PageShield = page_shield.NewPageShieldService(opts...)
	r.Rulesets = rulesets.NewRulesetService(opts...)
	r.URLNormalization = url_normalization.NewURLNormalizationService(opts...)
	r.Spectrum = spectrum.NewSpectrumService(opts...)
	r.Addressing = addressing.NewAddressingService(opts...)
	r.AuditLogs = audit_logs.NewAuditLogService(opts...)
	r.Billing = billing.NewBillingService(opts...)
	r.BrandProtection = brand_protection.NewBrandProtectionService(opts...)
	r.Diagnostics = diagnostics.NewDiagnosticService(opts...)
	r.Images = images.NewImageService(opts...)
	r.Intel = intel.NewIntelService(opts...)
	r.MagicTransit = magic_transit.NewMagicTransitService(opts...)
	r.MagicNetworkMonitoring = magic_network_monitoring.NewMagicNetworkMonitoringService(opts...)
	r.MTLSCertificates = mtls_certificates.NewMTLSCertificateService(opts...)
	r.Pages = pages.NewPageService(opts...)
	r.PCAPs = pcaps.NewPCAPService(opts...)
	r.Registrar = registrar.NewRegistrarService(opts...)
	r.RequestTracers = request_tracers.NewRequestTracerService(opts...)
	r.Rules = rules.NewRuleService(opts...)
	r.Storage = storage.NewStorageService(opts...)
	r.Stream = stream.NewStreamService(opts...)
	r.Alerting = alerting.NewAlertingService(opts...)
	r.D1 = d1.NewD1Service(opts...)
	r.R2 = r2.NewR2Service(opts...)
	r.WARPConnector = warp_connector.NewWARPConnectorService(opts...)
	r.WorkersForPlatforms = workers_for_platforms.NewWorkersForPlatformService(opts...)
	r.ZeroTrust = zero_trust.NewZeroTrustService(opts...)
	r.Challenges = challenges.NewChallengeService(opts...)
	r.Hyperdrive = hyperdrive.NewHyperdriveService(opts...)
	r.RUM = rum.NewRUMService(opts...)
	r.Vectorize = vectorize.NewVectorizeService(opts...)
	r.URLScanner = url_scanner.NewURLScannerService(opts...)
	r.Radar = radar.NewRadarService(opts...)
	r.BotManagement = bot_management.NewBotManagementService(opts...)
	r.OriginPostQuantumEncryption = origin_post_quantum_encryption.NewOriginPostQuantumEncryptionService(opts...)
	r.Speed = speed.NewSpeedService(opts...)
	r.DCVDelegation = dcv_delegation.NewDCVDelegationService(opts...)
	r.Hostnames = hostnames.NewHostnameService(opts...)
	r.Snippets = snippets.NewSnippetService(opts...)
	r.Calls = calls.NewCallService(opts...)
	r.CloudforceOne = cloudforce_one.NewCloudforceOneService(opts...)

	return
}
