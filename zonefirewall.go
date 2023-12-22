// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneFirewallService] method
// instead.
type ZoneFirewallService struct {
	Options     []option.RequestOption
	Lockdowns   *ZoneFirewallLockdownService
	Rules       *ZoneFirewallRuleService
	UaRules     *ZoneFirewallUaRuleService
	Wafs        *ZoneFirewallWafService
	AccessRules *ZoneFirewallAccessRuleService
}

// NewZoneFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneFirewallService(opts ...option.RequestOption) (r *ZoneFirewallService) {
	r = &ZoneFirewallService{}
	r.Options = opts
	r.Lockdowns = NewZoneFirewallLockdownService(opts...)
	r.Rules = NewZoneFirewallRuleService(opts...)
	r.UaRules = NewZoneFirewallUaRuleService(opts...)
	r.Wafs = NewZoneFirewallWafService(opts...)
	r.AccessRules = NewZoneFirewallAccessRuleService(opts...)
	return
}
