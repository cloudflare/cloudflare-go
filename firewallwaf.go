// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallWAFService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallWAFService] method
// instead.
type FirewallWAFService struct {
	Options   []option.RequestOption
	Overrides *FirewallWAFOverrideService
	Packages  *FirewallWAFPackageService
}

// NewFirewallWAFService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallWAFService(opts ...option.RequestOption) (r *FirewallWAFService) {
	r = &FirewallWAFService{}
	r.Options = opts
	r.Overrides = NewFirewallWAFOverrideService(opts...)
	r.Packages = NewFirewallWAFPackageService(opts...)
	return
}
