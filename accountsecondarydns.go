// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountSecondaryDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountSecondaryDNSService]
// method instead.
type AccountSecondaryDNSService struct {
	Options []option.RequestOption
	ACLs    *AccountSecondaryDNSACLService
	Peers   *AccountSecondaryDNSPeerService
	Tsigs   *AccountSecondaryDNSTsigService
}

// NewAccountSecondaryDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSService(opts ...option.RequestOption) (r *AccountSecondaryDNSService) {
	r = &AccountSecondaryDNSService{}
	r.Options = opts
	r.ACLs = NewAccountSecondaryDNSACLService(opts...)
	r.Peers = NewAccountSecondaryDNSPeerService(opts...)
	r.Tsigs = NewAccountSecondaryDNSTsigService(opts...)
	return
}
