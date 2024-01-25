// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountFirewallService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountFirewallService] method
// instead.
type AccountFirewallService struct {
	Options     []option.RequestOption
	AccessRules *AccountFirewallAccessRuleService
}

// NewAccountFirewallService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountFirewallService(opts ...option.RequestOption) (r *AccountFirewallService) {
	r = &AccountFirewallService{}
	r.Options = opts
	r.AccessRules = NewAccountFirewallAccessRuleService(opts...)
	return
}
