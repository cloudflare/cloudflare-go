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
	IPs                          *IPService
	Zones                        *ZoneService
	AI                           *AIService
	LoadBalancers                *LoadBalancerService
	Access                       *AccessService
	DNSRecords                   *DNSRecordService
	Emails                       *EmailService
	AccountMembers               *AccountMemberService
	Rules                        *RuleService
	Tunnels                      *TunnelService
	D1                           *D1Service
	Dex                          *DexService
	R2                           *R2Service
	Stream                       *StreamService
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
	VectorizeIndexes             *VectorizeIndexService
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
	Logpush                      *LogpushService
	Hold                         *HoldService
	PageShield                   *PageShieldService
	FontSettings                 *FontSettingService
	Snippets                     *SnippetService
	Dlp                          *DlpService
	Gateway                      *GatewayService
	AccessTags                   *AccessTagService
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_KEY, CLOUDFLARE_EMAIL, CLOUDFLARE_API_TOKEN). The
// option passed in as arguments are applied after these default arguments, and all
// option will be passed down to the services and requests that this client makes.
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
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.IPs = NewIPService(opts...)
	r.Zones = NewZoneService(opts...)
	r.AI = NewAIService(opts...)
	r.LoadBalancers = NewLoadBalancerService(opts...)
	r.Access = NewAccessService(opts...)
	r.DNSRecords = NewDNSRecordService(opts...)
	r.Emails = NewEmailService(opts...)
	r.AccountMembers = NewAccountMemberService(opts...)
	r.Rules = NewRuleService(opts...)
	r.Tunnels = NewTunnelService(opts...)
	r.D1 = NewD1Service(opts...)
	r.Dex = NewDexService(opts...)
	r.R2 = NewR2Service(opts...)
	r.Stream = NewStreamService(opts...)
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
	r.VectorizeIndexes = NewVectorizeIndexService(opts...)
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
	r.Logpush = NewLogpushService(opts...)
	r.Hold = NewHoldService(opts...)
	r.PageShield = NewPageShieldService(opts...)
	r.FontSettings = NewFontSettingService(opts...)
	r.Snippets = NewSnippetService(opts...)
	r.Dlp = NewDlpService(opts...)
	r.Gateway = NewGatewayService(opts...)
	r.AccessTags = NewAccessTagService(opts...)

	return
}
