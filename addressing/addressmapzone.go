// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AddressMapZoneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressMapZoneService] method instead.
type AddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAddressMapZoneService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressMapZoneService(opts ...option.RequestOption) (r *AddressMapZoneService) {
	r = &AddressMapZoneService{}
	r.Options = opts
	return
}
