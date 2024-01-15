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
	Options           []option.RequestOption
	Accounts          *AccountService
	Certificates      *CertificateService
	IPs               *IPService
	Memberships       *MembershipService
	Organizations     *OrganizationService
	Radar             *RadarService
	Railguns          *RailgunService
	User              *UserService
	Zones             *ZoneService
	AI                *AIService
	Alerting          *AlertingService
	Addressing        *AddressingService
	HyperdriveConfigs *HyperdriveConfigService
	ImagesV2s         *ImagesV2Service
	Intel             *IntelService
	WebAnalytics      *WebAnalyticService
	Stream            *StreamService
	Filters           *FilterService
	FirewallRules     *FirewallRuleService
	Healthchecks      *HealthcheckService
	Hostnames         *HostnameService
	Pagerules         *PageruleService
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_EMAIL, CLOUDFLARE_API_KEY, CLOUDFLARE_API_TOKEN,
// CLOUDFLARE_USER_SERVICE_KEY). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_EMAIL"); ok {
		defaults = append(defaults, option.WithAPIEmail(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_TOKEN"); ok {
		defaults = append(defaults, option.WithAPIToken(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_USER_SERVICE_KEY"); ok {
		defaults = append(defaults, option.WithUserServiceKey(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Accounts = NewAccountService(opts...)
	r.Certificates = NewCertificateService(opts...)
	r.IPs = NewIPService(opts...)
	r.Memberships = NewMembershipService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.Radar = NewRadarService(opts...)
	r.Railguns = NewRailgunService(opts...)
	r.User = NewUserService(opts...)
	r.Zones = NewZoneService(opts...)
	r.AI = NewAIService(opts...)
	r.Alerting = NewAlertingService(opts...)
	r.Addressing = NewAddressingService(opts...)
	r.HyperdriveConfigs = NewHyperdriveConfigService(opts...)
	r.ImagesV2s = NewImagesV2Service(opts...)
	r.Intel = NewIntelService(opts...)
	r.WebAnalytics = NewWebAnalyticService(opts...)
	r.Stream = NewStreamService(opts...)
	r.Filters = NewFilterService(opts...)
	r.FirewallRules = NewFirewallRuleService(opts...)
	r.Healthchecks = NewHealthcheckService(opts...)
	r.Hostnames = NewHostnameService(opts...)
	r.Pagerules = NewPageruleService(opts...)

	return
}
