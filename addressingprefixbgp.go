// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressingPrefixBgpService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingPrefixBgpService]
// method instead.
type AddressingPrefixBgpService struct {
	Options  []option.RequestOption
	Prefixes *AddressingPrefixBgpPrefixService
}

// NewAddressingPrefixBgpService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingPrefixBgpService(opts ...option.RequestOption) (r *AddressingPrefixBgpService) {
	r = &AddressingPrefixBgpService{}
	r.Options = opts
	r.Prefixes = NewAddressingPrefixBgpPrefixService(opts...)
	return
}
