// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressAddressMapService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressAddressMapService] method
// instead.
type AddressAddressMapService struct {
	Options  []option.RequestOption
	Accounts *AddressAddressMapAccountService
	IPs      *AddressAddressMapIPService
	Zones    *AddressAddressMapZoneService
}

// NewAddressAddressMapService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressAddressMapService(opts ...option.RequestOption) (r *AddressAddressMapService) {
	r = &AddressAddressMapService{}
	r.Options = opts
	r.Accounts = NewAddressAddressMapAccountService(opts...)
	r.IPs = NewAddressAddressMapIPService(opts...)
	r.Zones = NewAddressAddressMapZoneService(opts...)
	return
}
