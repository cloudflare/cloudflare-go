// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserFirewallService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserFirewallService] method
// instead.
type UserFirewallService struct {
	Options     []option.RequestOption
	AccessRules *UserFirewallAccessRuleService
}

// NewUserFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserFirewallService(opts ...option.RequestOption) (r *UserFirewallService) {
	r = &UserFirewallService{}
	r.Options = opts
	r.AccessRules = NewUserFirewallAccessRuleService(opts...)
	return
}
