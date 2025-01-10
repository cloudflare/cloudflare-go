// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go/v4/abuse_reports"
	"github.com/cloudflare/cloudflare-go/v4/accounts"
	"github.com/cloudflare/cloudflare-go/v4/acm"
	"github.com/cloudflare/cloudflare-go/v4/addressing"
	"github.com/cloudflare/cloudflare-go/v4/ai"
	"github.com/cloudflare/cloudflare-go/v4/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v4/alerting"
	"github.com/cloudflare/cloudflare-go/v4/api_gateway"
	"github.com/cloudflare/cloudflare-go/v4/argo"
	"github.com/cloudflare/cloudflare-go/v4/audit_logs"
	"github.com/cloudflare/cloudflare-go/v4/billing"
	"github.com/cloudflare/cloudflare-go/v4/bot_management"
	"github.com/cloudflare/cloudflare-go/v4/botnet_feed"
	"github.com/cloudflare/cloudflare-go/v4/brand_protection"
	"github.com/cloudflare/cloudflare-go/v4/cache"
	"github.com/cloudflare/cloudflare-go/v4/calls"
	"github.com/cloudflare/cloudflare-go/v4/certificate_authorities"
	"github.com/cloudflare/cloudflare-go/v4/client_certificates"
	"github.com/cloudflare/cloudflare-go/v4/cloud_connector"
	"github.com/cloudflare/cloudflare-go/v4/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v4/content_scanning"
	"github.com/cloudflare/cloudflare-go/v4/custom_certificates"
	"github.com/cloudflare/cloudflare-go/v4/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v4/custom_nameservers"
	"github.com/cloudflare/cloudflare-go/v4/d1"
	"github.com/cloudflare/cloudflare-go/v4/dcv_delegation"
	"github.com/cloudflare/cloudflare-go/v4/diagnostics"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/cloudflare/cloudflare-go/v4/dns_firewall"
	"github.com/cloudflare/cloudflare-go/v4/durable_objects"
	"github.com/cloudflare/cloudflare-go/v4/email_routing"
	"github.com/cloudflare/cloudflare-go/v4/email_security"
	"github.com/cloudflare/cloudflare-go/v4/filters"
	"github.com/cloudflare/cloudflare-go/v4/firewall"
	"github.com/cloudflare/cloudflare-go/v4/healthchecks"
	"github.com/cloudflare/cloudflare-go/v4/hostnames"
	"github.com/cloudflare/cloudflare-go/v4/hyperdrive"
	"github.com/cloudflare/cloudflare-go/v4/iam"
	"github.com/cloudflare/cloudflare-go/v4/images"
	"github.com/cloudflare/cloudflare-go/v4/intel"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/ips"
	"github.com/cloudflare/cloudflare-go/v4/keyless_certificates"
	"github.com/cloudflare/cloudflare-go/v4/kv"
	"github.com/cloudflare/cloudflare-go/v4/leaked_credential_checks"
	"github.com/cloudflare/cloudflare-go/v4/load_balancers"
	"github.com/cloudflare/cloudflare-go/v4/logpush"
	"github.com/cloudflare/cloudflare-go/v4/logs"
	"github.com/cloudflare/cloudflare-go/v4/magic_network_monitoring"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/cloudflare-go/v4/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v4/memberships"
	"github.com/cloudflare/cloudflare-go/v4/mtls_certificates"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/origin_ca_certificates"
	"github.com/cloudflare/cloudflare-go/v4/origin_post_quantum_encryption"
	"github.com/cloudflare/cloudflare-go/v4/origin_tls_client_auth"
	"github.com/cloudflare/cloudflare-go/v4/page_rules"
	"github.com/cloudflare/cloudflare-go/v4/page_shield"
	"github.com/cloudflare/cloudflare-go/v4/pages"
	"github.com/cloudflare/cloudflare-go/v4/queues"
	"github.com/cloudflare/cloudflare-go/v4/r2"
	"github.com/cloudflare/cloudflare-go/v4/radar"
	"github.com/cloudflare/cloudflare-go/v4/rate_limits"
	"github.com/cloudflare/cloudflare-go/v4/registrar"
	"github.com/cloudflare/cloudflare-go/v4/request_tracers"
	"github.com/cloudflare/cloudflare-go/v4/resource_sharing"
	"github.com/cloudflare/cloudflare-go/v4/rules"
	"github.com/cloudflare/cloudflare-go/v4/rulesets"
	"github.com/cloudflare/cloudflare-go/v4/rum"
	"github.com/cloudflare/cloudflare-go/v4/security_txt"
	"github.com/cloudflare/cloudflare-go/v4/snippets"
	"github.com/cloudflare/cloudflare-go/v4/spectrum"
	"github.com/cloudflare/cloudflare-go/v4/speed"
	"github.com/cloudflare/cloudflare-go/v4/ssl"
	"github.com/cloudflare/cloudflare-go/v4/stream"
	"github.com/cloudflare/cloudflare-go/v4/turnstile"
	"github.com/cloudflare/cloudflare-go/v4/url_normalization"
	"github.com/cloudflare/cloudflare-go/v4/url_scanner"
	"github.com/cloudflare/cloudflare-go/v4/user"
	"github.com/cloudflare/cloudflare-go/v4/vectorize"
	"github.com/cloudflare/cloudflare-go/v4/waiting_rooms"
	"github.com/cloudflare/cloudflare-go/v4/web3"
	"github.com/cloudflare/cloudflare-go/v4/workers"
	"github.com/cloudflare/cloudflare-go/v4/workers_for_platforms"
	"github.com/cloudflare/cloudflare-go/v4/workflows"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/cloudflare-go/v4/zones"
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
	ACM                         *acm.ACMService
	Argo                        *argo.ArgoService
	CertificateAuthorities      *certificate_authorities.CertificateAuthorityService
	ClientCertificates          *client_certificates.ClientCertificateService
	CustomCertificates          *custom_certificates.CustomCertificateService
	CustomHostnames             *custom_hostnames.CustomHostnameService
	CustomNameservers           *custom_nameservers.CustomNameserverService
	DNSFirewall                 *dns_firewall.DNSFirewallService
	DNS                         *dns.DNSService
	EmailSecurity               *email_security.EmailSecurityService
	EmailRouting                *email_routing.EmailRoutingService
	Filters                     *filters.FilterService
	Firewall                    *firewall.FirewallService
	Healthchecks                *healthchecks.HealthcheckService
	KeylessCertificates         *keyless_certificates.KeylessCertificateService
	Logpush                     *logpush.LogpushService
	Logs                        *logs.LogService
	OriginTLSClientAuth         *origin_tls_client_auth.OriginTLSClientAuthService
	PageRules                   *page_rules.PageRuleService
	RateLimits                  *rate_limits.RateLimitService
	WaitingRooms                *waiting_rooms.WaitingRoomService
	Web3                        *web3.Web3Service
	Workers                     *workers.WorkerService
	KV                          *kv.KVService
	DurableObjects              *durable_objects.DurableObjectService
	Queues                      *queues.QueueService
	APIGateway                  *api_gateway.APIGatewayService
	ManagedTransforms           *managed_transforms.ManagedTransformService
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
	Registrar                   *registrar.RegistrarService
	RequestTracers              *request_tracers.RequestTracerService
	Rules                       *rules.RuleService
	Stream                      *stream.StreamService
	Alerting                    *alerting.AlertingService
	D1                          *d1.D1Service
	R2                          *r2.R2Service
	WorkersForPlatforms         *workers_for_platforms.WorkersForPlatformService
	ZeroTrust                   *zero_trust.ZeroTrustService
	Turnstile                   *turnstile.TurnstileService
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
	AIGateway                   *ai_gateway.AIGatewayService
	IAM                         *iam.IAMService
	CloudConnector              *cloud_connector.CloudConnectorService
	BotnetFeed                  *botnet_feed.BotnetFeedService
	SecurityTXT                 *security_txt.SecurityTXTService
	Workflows                   *workflows.WorkflowService
	ResourceSharing             *resource_sharing.ResourceSharingService
	LeakedCredentialChecks      *leaked_credential_checks.LeakedCredentialCheckService
	ContentScanning             *content_scanning.ContentScanningService
	AbuseReports                *abuse_reports.AbuseReportService
	AI                          *ai.AIService
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
	r.ACM = acm.NewACMService(opts...)
	r.Argo = argo.NewArgoService(opts...)
	r.CertificateAuthorities = certificate_authorities.NewCertificateAuthorityService(opts...)
	r.ClientCertificates = client_certificates.NewClientCertificateService(opts...)
	r.CustomCertificates = custom_certificates.NewCustomCertificateService(opts...)
	r.CustomHostnames = custom_hostnames.NewCustomHostnameService(opts...)
	r.CustomNameservers = custom_nameservers.NewCustomNameserverService(opts...)
	r.DNSFirewall = dns_firewall.NewDNSFirewallService(opts...)
	r.DNS = dns.NewDNSService(opts...)
	r.EmailSecurity = email_security.NewEmailSecurityService(opts...)
	r.EmailRouting = email_routing.NewEmailRoutingService(opts...)
	r.Filters = filters.NewFilterService(opts...)
	r.Firewall = firewall.NewFirewallService(opts...)
	r.Healthchecks = healthchecks.NewHealthcheckService(opts...)
	r.KeylessCertificates = keyless_certificates.NewKeylessCertificateService(opts...)
	r.Logpush = logpush.NewLogpushService(opts...)
	r.Logs = logs.NewLogService(opts...)
	r.OriginTLSClientAuth = origin_tls_client_auth.NewOriginTLSClientAuthService(opts...)
	r.PageRules = page_rules.NewPageRuleService(opts...)
	r.RateLimits = rate_limits.NewRateLimitService(opts...)
	r.WaitingRooms = waiting_rooms.NewWaitingRoomService(opts...)
	r.Web3 = web3.NewWeb3Service(opts...)
	r.Workers = workers.NewWorkerService(opts...)
	r.KV = kv.NewKVService(opts...)
	r.DurableObjects = durable_objects.NewDurableObjectService(opts...)
	r.Queues = queues.NewQueueService(opts...)
	r.APIGateway = api_gateway.NewAPIGatewayService(opts...)
	r.ManagedTransforms = managed_transforms.NewManagedTransformService(opts...)
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
	r.Registrar = registrar.NewRegistrarService(opts...)
	r.RequestTracers = request_tracers.NewRequestTracerService(opts...)
	r.Rules = rules.NewRuleService(opts...)
	r.Stream = stream.NewStreamService(opts...)
	r.Alerting = alerting.NewAlertingService(opts...)
	r.D1 = d1.NewD1Service(opts...)
	r.R2 = r2.NewR2Service(opts...)
	r.WorkersForPlatforms = workers_for_platforms.NewWorkersForPlatformService(opts...)
	r.ZeroTrust = zero_trust.NewZeroTrustService(opts...)
	r.Turnstile = turnstile.NewTurnstileService(opts...)
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
	r.AIGateway = ai_gateway.NewAIGatewayService(opts...)
	r.IAM = iam.NewIAMService(opts...)
	r.CloudConnector = cloud_connector.NewCloudConnectorService(opts...)
	r.BotnetFeed = botnet_feed.NewBotnetFeedService(opts...)
	r.SecurityTXT = security_txt.NewSecurityTXTService(opts...)
	r.Workflows = workflows.NewWorkflowService(opts...)
	r.ResourceSharing = resource_sharing.NewResourceSharingService(opts...)
	r.LeakedCredentialChecks = leaked_credential_checks.NewLeakedCredentialCheckService(opts...)
	r.ContentScanning = content_scanning.NewContentScanningService(opts...)
	r.AbuseReports = abuse_reports.NewAbuseReportService(opts...)
	r.AI = ai.NewAIService(opts...)

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
