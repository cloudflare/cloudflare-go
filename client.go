// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/cloudflare/cloudflare-go/v7/abuse_reports"
	"github.com/cloudflare/cloudflare-go/v7/accounts"
	"github.com/cloudflare/cloudflare-go/v7/acm"
	"github.com/cloudflare/cloudflare-go/v7/addressing"
	"github.com/cloudflare/cloudflare-go/v7/ai"
	"github.com/cloudflare/cloudflare-go/v7/ai_audit"
	"github.com/cloudflare/cloudflare-go/v7/ai_gateway"
	"github.com/cloudflare/cloudflare-go/v7/ai_search"
	"github.com/cloudflare/cloudflare-go/v7/ai_security"
	"github.com/cloudflare/cloudflare-go/v7/alerting"
	"github.com/cloudflare/cloudflare-go/v7/api_gateway"
	"github.com/cloudflare/cloudflare-go/v7/argo"
	"github.com/cloudflare/cloudflare-go/v7/audit_logs"
	"github.com/cloudflare/cloudflare-go/v7/billing"
	"github.com/cloudflare/cloudflare-go/v7/bot_management"
	"github.com/cloudflare/cloudflare-go/v7/botnet_feed"
	"github.com/cloudflare/cloudflare-go/v7/brand_protection"
	"github.com/cloudflare/cloudflare-go/v7/browser_rendering"
	"github.com/cloudflare/cloudflare-go/v7/cache"
	"github.com/cloudflare/cloudflare-go/v7/calls"
	"github.com/cloudflare/cloudflare-go/v7/certificate_authorities"
	"github.com/cloudflare/cloudflare-go/v7/client_certificates"
	"github.com/cloudflare/cloudflare-go/v7/cloud_connector"
	"github.com/cloudflare/cloudflare-go/v7/cloudforce_one"
	"github.com/cloudflare/cloudflare-go/v7/connectivity"
	"github.com/cloudflare/cloudflare-go/v7/content_scanning"
	"github.com/cloudflare/cloudflare-go/v7/custom_certificates"
	"github.com/cloudflare/cloudflare-go/v7/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v7/custom_nameservers"
	"github.com/cloudflare/cloudflare-go/v7/custom_pages"
	"github.com/cloudflare/cloudflare-go/v7/d1"
	"github.com/cloudflare/cloudflare-go/v7/dcv_delegation"
	"github.com/cloudflare/cloudflare-go/v7/ddos_protection"
	"github.com/cloudflare/cloudflare-go/v7/diagnostics"
	"github.com/cloudflare/cloudflare-go/v7/dns"
	"github.com/cloudflare/cloudflare-go/v7/dns_firewall"
	"github.com/cloudflare/cloudflare-go/v7/durable_objects"
	"github.com/cloudflare/cloudflare-go/v7/email_routing"
	"github.com/cloudflare/cloudflare-go/v7/email_security"
	"github.com/cloudflare/cloudflare-go/v7/email_sending"
	"github.com/cloudflare/cloudflare-go/v7/filters"
	"github.com/cloudflare/cloudflare-go/v7/firewall"
	"github.com/cloudflare/cloudflare-go/v7/fraud"
	"github.com/cloudflare/cloudflare-go/v7/google_tag_gateway"
	"github.com/cloudflare/cloudflare-go/v7/healthchecks"
	"github.com/cloudflare/cloudflare-go/v7/hostnames"
	"github.com/cloudflare/cloudflare-go/v7/hyperdrive"
	"github.com/cloudflare/cloudflare-go/v7/iam"
	"github.com/cloudflare/cloudflare-go/v7/images"
	"github.com/cloudflare/cloudflare-go/v7/intel"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/ips"
	"github.com/cloudflare/cloudflare-go/v7/keyless_certificates"
	"github.com/cloudflare/cloudflare-go/v7/kv"
	"github.com/cloudflare/cloudflare-go/v7/leaked_credential_checks"
	"github.com/cloudflare/cloudflare-go/v7/load_balancers"
	"github.com/cloudflare/cloudflare-go/v7/logpush"
	"github.com/cloudflare/cloudflare-go/v7/logs"
	"github.com/cloudflare/cloudflare-go/v7/magic_cloud_networking"
	"github.com/cloudflare/cloudflare-go/v7/magic_network_monitoring"
	"github.com/cloudflare/cloudflare-go/v7/magic_transit"
	"github.com/cloudflare/cloudflare-go/v7/managed_transforms"
	"github.com/cloudflare/cloudflare-go/v7/memberships"
	"github.com/cloudflare/cloudflare-go/v7/mtls_certificates"
	"github.com/cloudflare/cloudflare-go/v7/network_interconnects"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/organizations"
	"github.com/cloudflare/cloudflare-go/v7/origin_ca_certificates"
	"github.com/cloudflare/cloudflare-go/v7/origin_post_quantum_encryption"
	"github.com/cloudflare/cloudflare-go/v7/origin_tls_client_auth"
	"github.com/cloudflare/cloudflare-go/v7/page_rules"
	"github.com/cloudflare/cloudflare-go/v7/page_shield"
	"github.com/cloudflare/cloudflare-go/v7/pages"
	"github.com/cloudflare/cloudflare-go/v7/pipelines"
	"github.com/cloudflare/cloudflare-go/v7/queues"
	"github.com/cloudflare/cloudflare-go/v7/r2"
	"github.com/cloudflare/cloudflare-go/v7/r2_data_catalog"
	"github.com/cloudflare/cloudflare-go/v7/radar"
	"github.com/cloudflare/cloudflare-go/v7/rate_limits"
	"github.com/cloudflare/cloudflare-go/v7/realtime_kit"
	"github.com/cloudflare/cloudflare-go/v7/registrar"
	"github.com/cloudflare/cloudflare-go/v7/request_tracers"
	"github.com/cloudflare/cloudflare-go/v7/resource_sharing"
	"github.com/cloudflare/cloudflare-go/v7/resource_tagging"
	"github.com/cloudflare/cloudflare-go/v7/rules"
	"github.com/cloudflare/cloudflare-go/v7/rulesets"
	"github.com/cloudflare/cloudflare-go/v7/rum"
	"github.com/cloudflare/cloudflare-go/v7/schema_validation"
	"github.com/cloudflare/cloudflare-go/v7/secrets_store"
	"github.com/cloudflare/cloudflare-go/v7/security_center"
	"github.com/cloudflare/cloudflare-go/v7/security_txt"
	"github.com/cloudflare/cloudflare-go/v7/snippets"
	"github.com/cloudflare/cloudflare-go/v7/spectrum"
	"github.com/cloudflare/cloudflare-go/v7/speed"
	"github.com/cloudflare/cloudflare-go/v7/ssl"
	"github.com/cloudflare/cloudflare-go/v7/stream"
	"github.com/cloudflare/cloudflare-go/v7/token_validation"
	"github.com/cloudflare/cloudflare-go/v7/turnstile"
	"github.com/cloudflare/cloudflare-go/v7/url_normalization"
	"github.com/cloudflare/cloudflare-go/v7/url_scanner"
	"github.com/cloudflare/cloudflare-go/v7/user"
	"github.com/cloudflare/cloudflare-go/v7/vectorize"
	"github.com/cloudflare/cloudflare-go/v7/vulnerability_scanner"
	"github.com/cloudflare/cloudflare-go/v7/waiting_rooms"
	"github.com/cloudflare/cloudflare-go/v7/web3"
	"github.com/cloudflare/cloudflare-go/v7/workers"
	"github.com/cloudflare/cloudflare-go/v7/workers_for_platforms"
	"github.com/cloudflare/cloudflare-go/v7/workflows"
	"github.com/cloudflare/cloudflare-go/v7/zaraz"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/cloudflare-go/v7/zones"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                []option.RequestOption
	Accounts               *accounts.AccountService
	Organizations          *organizations.OrganizationService
	OriginCACertificates   *origin_ca_certificates.OriginCACertificateService
	IPs                    *ips.IPService
	Memberships            *memberships.MembershipService
	User                   *user.UserService
	Zones                  *zones.ZoneService
	LoadBalancers          *load_balancers.LoadBalancerService
	Cache                  *cache.CacheService
	SSL                    *ssl.SSLService
	ACM                    *acm.ACMService
	Argo                   *argo.ArgoService
	CertificateAuthorities *certificate_authorities.CertificateAuthorityService
	ClientCertificates     *client_certificates.ClientCertificateService
	CustomCertificates     *custom_certificates.CustomCertificateService
	CustomHostnames        *custom_hostnames.CustomHostnameService
	CustomNameservers      *custom_nameservers.CustomNameserverService
	DNSFirewall            *dns_firewall.DNSFirewallService
	DNS                    *dns.DNSService
	EmailSecurity          *email_security.EmailSecurityService
	EmailRouting           *email_routing.EmailRoutingService
	EmailSending           *email_sending.EmailSendingService
	// Deprecated: The Filters API is deprecated in favour of using the Ruleset Engine.
	// See
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
	// for full details.
	Filters             *filters.FilterService
	Firewall            *firewall.FirewallService
	Healthchecks        *healthchecks.HealthcheckService
	KeylessCertificates *keyless_certificates.KeylessCertificateService
	Logpush             *logpush.LogpushService
	Logs                *logs.LogService
	OriginTLSClientAuth *origin_tls_client_auth.OriginTLSClientAuthService
	PageRules           *page_rules.PageRuleService
	// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
	// Engine. See
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
	// for full details.
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
	DDoSProtection              *ddos_protection.DDoSProtectionService
	MagicNetworkMonitoring      *magic_network_monitoring.MagicNetworkMonitoringService
	MagicCloudNetworking        *magic_cloud_networking.MagicCloudNetworkingService
	NetworkInterconnects        *network_interconnects.NetworkInterconnectService
	MTLSCertificates            *mtls_certificates.MTLSCertificateService
	Pages                       *pages.PageService
	Registrar                   *registrar.RegistrarService
	RequestTracers              *request_tracers.RequestTracerService
	Rules                       *rules.RuleService
	Stream                      *stream.StreamService
	Alerting                    *alerting.AlertingService
	D1                          *d1.D1Service
	R2                          *r2.R2Service
	R2DataCatalog               *r2_data_catalog.R2DataCatalogService
	WorkersForPlatforms         *workers_for_platforms.WorkersForPlatformService
	ZeroTrust                   *zero_trust.ZeroTrustService
	Turnstile                   *turnstile.TurnstileService
	Connectivity                *connectivity.ConnectivityService
	Hyperdrive                  *hyperdrive.HyperdriveService
	RUM                         *rum.RUMService
	Vectorize                   *vectorize.VectorizeService
	URLScanner                  *url_scanner.URLScannerService
	VulnerabilityScanner        *vulnerability_scanner.VulnerabilityScannerService
	Radar                       *radar.RadarService
	BotManagement               *bot_management.BotManagementService
	Fraud                       *fraud.FraudService
	OriginPostQuantumEncryption *origin_post_quantum_encryption.OriginPostQuantumEncryptionService
	GoogleTagGateway            *google_tag_gateway.GoogleTagGatewayService
	Zaraz                       *zaraz.ZarazService
	Speed                       *speed.SpeedService
	DCVDelegation               *dcv_delegation.DCVDelegationService
	Hostnames                   *hostnames.HostnameService
	Snippets                    *snippets.SnippetService
	RealtimeKit                 *realtime_kit.RealtimeKitService
	Calls                       *calls.CallService
	CloudforceOne               *cloudforce_one.CloudforceOneService
	AIGateway                   *ai_gateway.AIGatewayService
	IAM                         *iam.IAMService
	CloudConnector              *cloud_connector.CloudConnectorService
	BotnetFeed                  *botnet_feed.BotnetFeedService
	SecurityTXT                 *security_txt.SecurityTXTService
	Workflows                   *workflows.WorkflowService
	ResourceSharing             *resource_sharing.ResourceSharingService
	ResourceTagging             *resource_tagging.ResourceTaggingService
	LeakedCredentialChecks      *leaked_credential_checks.LeakedCredentialCheckService
	ContentScanning             *content_scanning.ContentScanningService
	AISecurity                  *ai_security.AISecurityService
	AbuseReports                *abuse_reports.AbuseReportService
	AI                          *ai.AIService
	AIAudit                     *ai_audit.AIAuditService
	AISearch                    *ai_search.AISearchService
	SecurityCenter              *security_center.SecurityCenterService
	BrowserRendering            *browser_rendering.BrowserRenderingService
	CustomPages                 *custom_pages.CustomPageService
	SecretsStore                *secrets_store.SecretsStoreService
	Pipelines                   *pipelines.PipelineService
	SchemaValidation            *schema_validation.SchemaValidationService
	TokenValidation             *token_validation.TokenValidationService
}

// DefaultClientOptions read from the environment (CLOUDFLARE_API_KEY,
// CLOUDFLARE_API_USER_SERVICE_KEY, CLOUDFLARE_API_TOKEN, CLOUDFLARE_EMAIL,
// CLOUDFLARE_BASE_URL). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithHTTPClient(defaultHTTPClient()), option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
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
	if o, ok := os.LookupEnv("CLOUDFLARE_CUSTOM_HEADERS"); ok {
		for _, line := range strings.Split(o, "\n") {
			colon := strings.Index(line, ":")
			if colon >= 0 {
				defaults = append(defaults, option.WithHeader(strings.TrimSpace(line[:colon]), strings.TrimSpace(line[colon+1:])))
			}
		}
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_KEY, CLOUDFLARE_API_USER_SERVICE_KEY,
// CLOUDFLARE_API_TOKEN, CLOUDFLARE_EMAIL, CLOUDFLARE_BASE_URL). The option passed
// in as arguments are applied after these default arguments, and all option will
// be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = &Client{Options: opts}

	r.Accounts = accounts.NewAccountService(opts...)
	r.Organizations = organizations.NewOrganizationService(opts...)
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
	r.EmailSending = email_sending.NewEmailSendingService(opts...)
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
	r.DDoSProtection = ddos_protection.NewDDoSProtectionService(opts...)
	r.MagicNetworkMonitoring = magic_network_monitoring.NewMagicNetworkMonitoringService(opts...)
	r.MagicCloudNetworking = magic_cloud_networking.NewMagicCloudNetworkingService(opts...)
	r.NetworkInterconnects = network_interconnects.NewNetworkInterconnectService(opts...)
	r.MTLSCertificates = mtls_certificates.NewMTLSCertificateService(opts...)
	r.Pages = pages.NewPageService(opts...)
	r.Registrar = registrar.NewRegistrarService(opts...)
	r.RequestTracers = request_tracers.NewRequestTracerService(opts...)
	r.Rules = rules.NewRuleService(opts...)
	r.Stream = stream.NewStreamService(opts...)
	r.Alerting = alerting.NewAlertingService(opts...)
	r.D1 = d1.NewD1Service(opts...)
	r.R2 = r2.NewR2Service(opts...)
	r.R2DataCatalog = r2_data_catalog.NewR2DataCatalogService(opts...)
	r.WorkersForPlatforms = workers_for_platforms.NewWorkersForPlatformService(opts...)
	r.ZeroTrust = zero_trust.NewZeroTrustService(opts...)
	r.Turnstile = turnstile.NewTurnstileService(opts...)
	r.Connectivity = connectivity.NewConnectivityService(opts...)
	r.Hyperdrive = hyperdrive.NewHyperdriveService(opts...)
	r.RUM = rum.NewRUMService(opts...)
	r.Vectorize = vectorize.NewVectorizeService(opts...)
	r.URLScanner = url_scanner.NewURLScannerService(opts...)
	r.VulnerabilityScanner = vulnerability_scanner.NewVulnerabilityScannerService(opts...)
	r.Radar = radar.NewRadarService(opts...)
	r.BotManagement = bot_management.NewBotManagementService(opts...)
	r.Fraud = fraud.NewFraudService(opts...)
	r.OriginPostQuantumEncryption = origin_post_quantum_encryption.NewOriginPostQuantumEncryptionService(opts...)
	r.GoogleTagGateway = google_tag_gateway.NewGoogleTagGatewayService(opts...)
	r.Zaraz = zaraz.NewZarazService(opts...)
	r.Speed = speed.NewSpeedService(opts...)
	r.DCVDelegation = dcv_delegation.NewDCVDelegationService(opts...)
	r.Hostnames = hostnames.NewHostnameService(opts...)
	r.Snippets = snippets.NewSnippetService(opts...)
	r.RealtimeKit = realtime_kit.NewRealtimeKitService(opts...)
	r.Calls = calls.NewCallService(opts...)
	r.CloudforceOne = cloudforce_one.NewCloudforceOneService(opts...)
	r.AIGateway = ai_gateway.NewAIGatewayService(opts...)
	r.IAM = iam.NewIAMService(opts...)
	r.CloudConnector = cloud_connector.NewCloudConnectorService(opts...)
	r.BotnetFeed = botnet_feed.NewBotnetFeedService(opts...)
	r.SecurityTXT = security_txt.NewSecurityTXTService(opts...)
	r.Workflows = workflows.NewWorkflowService(opts...)
	r.ResourceSharing = resource_sharing.NewResourceSharingService(opts...)
	r.ResourceTagging = resource_tagging.NewResourceTaggingService(opts...)
	r.LeakedCredentialChecks = leaked_credential_checks.NewLeakedCredentialCheckService(opts...)
	r.ContentScanning = content_scanning.NewContentScanningService(opts...)
	r.AISecurity = ai_security.NewAISecurityService(opts...)
	r.AbuseReports = abuse_reports.NewAbuseReportService(opts...)
	r.AI = ai.NewAIService(opts...)
	r.AIAudit = ai_audit.NewAIAuditService(opts...)
	r.AISearch = ai_search.NewAISearchService(opts...)
	r.SecurityCenter = security_center.NewSecurityCenterService(opts...)
	r.BrowserRendering = browser_rendering.NewBrowserRenderingService(opts...)
	r.CustomPages = custom_pages.NewCustomPageService(opts...)
	r.SecretsStore = secrets_store.NewSecretsStoreService(opts...)
	r.Pipelines = pipelines.NewPipelineService(opts...)
	r.SchemaValidation = schema_validation.NewSchemaValidationService(opts...)
	r.TokenValidation = token_validation.NewTokenValidationService(opts...)

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
	opts = slices.Concat(r.Options, opts)
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
