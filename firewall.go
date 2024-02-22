// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewFirewallService] method instead.
type FirewallService struct {
	Options     []option.RequestOption
	Lockdowns   *FirewallLockdownService
	Rules       *FirewallRuleService
	AccessRules *FirewallAccessRuleService
	UARules     *FirewallUARuleService
	WAF         *FirewallWAFService
}

// NewFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallService(opts ...option.RequestOption) (r *FirewallService) {
	r = &FirewallService{}
	r.Options = opts
	r.Lockdowns = NewFirewallLockdownService(opts...)
	r.Rules = NewFirewallRuleService(opts...)
	r.AccessRules = NewFirewallAccessRuleService(opts...)
	r.UARules = NewFirewallUARuleService(opts...)
	r.WAF = NewFirewallWAFService(opts...)
	return
}
