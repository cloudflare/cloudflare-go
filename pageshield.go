// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// PageShieldService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPageShieldService] method instead.
type PageShieldService struct {
	Options     []option.RequestOption
	Connections *PageShieldConnectionService
	Policies    *PageShieldPolicyService
}

// NewPageShieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageShieldService(opts ...option.RequestOption) (r *PageShieldService) {
	r = &PageShieldService{}
	r.Options = opts
	r.Connections = NewPageShieldConnectionService(opts...)
	r.Policies = NewPageShieldPolicyService(opts...)
	return
}
