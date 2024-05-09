// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go/v2/accounts"
	"github.com/cloudflare/cloudflare-go/v2/acm"
	"github.com/cloudflare/cloudflare-go/v2/addressing"
	"github.com/cloudflare/cloudflare-go/v2/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v2/alerting"
	"github.com/cloudflare/cloudflare-go/v2/argo"
	"github.com/cloudflare/cloudflare-go/v2/audit_logs"
	"github.com/cloudflare/cloudflare-go/v2/billing"
	"github.com/cloudflare/cloudflare-go/v2/bot_management"
	"github.com/cloudflare/cloudflare-go/v2/brand_protection"
	"github.com/cloudflare/cloudflare-go/v2/cache"
	"github.com/cloudflare/cloudflare-go/v2/calls"
	"github.com/cloudflare/cloudflare-go/v2/certificate_authorities"
	"github.com/cloudflare/cloudflare-go/v2/challenges"
	"github.com/cloudflare/cloudflare-go/v2/client_certificates"
	"github.com/cloudflare/cloudflare-go/v2/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v2/custom_certificates"
	"github.com/cloudflare/cloudflare-go/v2/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v2/custom_nameservers"
	"github.com/cloudflare/cloudflare-go/v2/d1"
	"github.com/cloudflare/cloudflare-go/v2/dcv_delegation"
	"github.com/cloudflare/cloudflare-go/v2/diagnostics"
	"github.com/cloudflare/cloudflare-go/v2/dns"
	"github.com/cloudflare/cloudflare-go/v2/dnssec"
	"github.com/cloudflare/cloudflare-go/v2/durable_objects"
	"github.com/cloudflare/cloudflare-go/v2/email_routing"
	"github.com/cloudflare/cloudflare-go/v2/event_notifications"
	"github.com/cloudflare/cloudflare-go/v2/filters"
	"github.com/cloudflare/cloudflare-go/v2/firewall"
	"github.com/cloudflare/cloudflare-go/v2/healthchecks"
	"github.com/cloudflare/cloudflare-go/v2/hostnames"
	"github.com/cloudflare/cloudflare-go/v2/hyperdrive"
	"github.com/cloudflare/cloudflare-go/v2/images"
	"github.com/cloudflare/cloudflare-go/v2/intel"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/ips"
	"github.com/cloudflare/cloudflare-go/v2/keyless_certificates"
	"github.com/cloudflare/cloudflare-go/v2/kv"
	"github.com/cloudflare/cloudflare-go/v2/load_balancers"
	"github.com/cloudflare/cloudflare-go/v2/logpush"
	"github.com/cloudflare/cloudflare-go/v2/logs"
	"github.com/cloudflare/cloudflare-go/v2/magic_network_monitoring"
	"github.com/cloudflare/cloudflare-go/v2/magic_transit"
	"github.com/cloudflare/cloudflare-go/v2/managed_headers"
	"github.com/cloudflare/cloudflare-go/v2/memberships"
	"github.com/cloudflare/cloudflare-go/v2/mtls_certificates"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/origin_ca_certificates"
	"github.com/cloudflare/cloudflare-go/v2/origin_post_quantum_encryption"
	"github.com/cloudflare/cloudflare-go/v2/origin_tls_client_auth"
	"github.com/cloudflare/cloudflare-go/v2/page_shield"
	"github.com/cloudflare/cloudflare-go/v2/pagerules"
	"github.com/cloudflare/cloudflare-go/v2/pages"
	"github.com/cloudflare/cloudflare-go/v2/pcaps"
	"github.com/cloudflare/cloudflare-go/v2/plans"
	"github.com/cloudflare/cloudflare-go/v2/queues"
	"github.com/cloudflare/cloudflare-go/v2/r2"
	"github.com/cloudflare/cloudflare-go/v2/radar"
	"github.com/cloudflare/cloudflare-go/v2/rate_limits"
	"github.com/cloudflare/cloudflare-go/v2/rate_plans"
	"github.com/cloudflare/cloudflare-go/v2/registrar"
	"github.com/cloudflare/cloudflare-go/v2/request_tracers"
	"github.com/cloudflare/cloudflare-go/v2/rules"
	"github.com/cloudflare/cloudflare-go/v2/rulesets"
	"github.com/cloudflare/cloudflare-go/v2/rum"
	"github.com/cloudflare/cloudflare-go/v2/secondary_dns"
	"github.com/cloudflare/cloudflare-go/v2/snippets"
	"github.com/cloudflare/cloudflare-go/v2/spectrum"
	"github.com/cloudflare/cloudflare-go/v2/speed"
	"github.com/cloudflare/cloudflare-go/v2/ssl"
	"github.com/cloudflare/cloudflare-go/v2/storage"
	"github.com/cloudflare/cloudflare-go/v2/stream"
	"github.com/cloudflare/cloudflare-go/v2/subscriptions"
	"github.com/cloudflare/cloudflare-go/v2/url_normalization"
	"github.com/cloudflare/cloudflare-go/v2/url_scanner"
	"github.com/cloudflare/cloudflare-go/v2/user"
	"github.com/cloudflare/cloudflare-go/v2/vectorize"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
	"github.com/cloudflare/cloudflare-go/v2/warp_connector"
	"github.com/cloudflare/cloudflare-go/v2/web3"
	"github.com/cloudflare/cloudflare-go/v2/workers"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/cloudflare-go/v2/zones"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                     []option.RequestOption
	Accounts                    *accounts.AccountService
	OriginCACertificates        *origin_ca_certificates.OriginCACertificateService
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
	Plans                       *plans.PlanService
	RatePlans                   *rate_plans.RatePlanService
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
	EventNotifications          *event_notifications.EventNotificationService
	AIGateway                   *ai_gateway.AIGatewayService
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_TOKEN, CLOUDFLARE_API_KEY, CLOUDFLARE_EMAIL,
// CLOUDFLARE_API_USER_SERVICE_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_TOKEN"); ok {
		defaults = append(defaults, option.WithAPIToken(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_EMAIL"); ok {
		defaults = append(defaults, option.WithAPIEmail(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_USER_SERVICE_KEY"); ok {
		defaults = append(defaults, option.WithUserServiceKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = accounts.NewAccountService(opts...)
	r.OriginCACertificates = origin_ca_certificates.NewOriginCACertificateService(opts...)
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
	r.Plans = plans.NewPlanService(opts...)
	r.RatePlans = rate_plans.NewRatePlanService(opts...)
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
	r.EventNotifications = event_notifications.NewEventNotificationService(opts...)
	r.AIGateway = ai_gateway.NewAIGatewayService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
