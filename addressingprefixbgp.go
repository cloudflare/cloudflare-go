// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// AddressingPrefixBGPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingPrefixBGPService]
// method instead.
type AddressingPrefixBGPService struct {
	Options  []option.RequestOption
	Bindings *AddressingPrefixBGPBindingService
	Prefixes *AddressingPrefixBGPPrefixService
	Statuses *AddressingPrefixBGPStatusService
}

// NewAddressingPrefixBGPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingPrefixBGPService(opts ...option.RequestOption) (r *AddressingPrefixBGPService) {
	r = &AddressingPrefixBGPService{}
	r.Options = opts
	r.Bindings = NewAddressingPrefixBGPBindingService(opts...)
	r.Prefixes = NewAddressingPrefixBGPPrefixService(opts...)
	r.Statuses = NewAddressingPrefixBGPStatusService(opts...)
	return
}
