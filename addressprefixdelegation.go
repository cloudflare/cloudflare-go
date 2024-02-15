// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressPrefixDelegationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressPrefixDelegationService] method instead.
type AddressPrefixDelegationService struct {
	Options []option.RequestOption
}

// NewAddressPrefixDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressPrefixDelegationService(opts ...option.RequestOption) (r *AddressPrefixDelegationService) {
	r = &AddressPrefixDelegationService{}
	r.Options = opts
	return
}
