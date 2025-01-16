// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// PrefixBGPService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixBGPService] method instead.
type PrefixBGPService struct {
	Options  []option.RequestOption
	Bindings *PrefixBGPBindingService
	Prefixes *PrefixBGPPrefixService
	Statuses *PrefixBGPStatusService
}

// NewPrefixBGPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrefixBGPService(opts ...option.RequestOption) (r *PrefixBGPService) {
	r = &PrefixBGPService{}
	r.Options = opts
	r.Bindings = NewPrefixBGPBindingService(opts...)
	r.Prefixes = NewPrefixBGPPrefixService(opts...)
	r.Statuses = NewPrefixBGPStatusService(opts...)
	return
}
