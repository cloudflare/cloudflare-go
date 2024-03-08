// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// SecondaryDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSService] method
// instead.
type SecondaryDNSService struct {
	Options    []option.RequestOption
	ForceAxfrs *SecondaryDNSForceAxfrService
	Incoming   *SecondaryDNSIncomingService
	Outgoing   *SecondaryDNSOutgoingService
	ACLs       *SecondaryDNSACLService
	Peers      *SecondaryDNSPeerService
	TSIGs      *SecondaryDNSTSIGService
}

// NewSecondaryDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecondaryDNSService(opts ...option.RequestOption) (r *SecondaryDNSService) {
	r = &SecondaryDNSService{}
	r.Options = opts
	r.ForceAxfrs = NewSecondaryDNSForceAxfrService(opts...)
	r.Incoming = NewSecondaryDNSIncomingService(opts...)
	r.Outgoing = NewSecondaryDNSOutgoingService(opts...)
	r.ACLs = NewSecondaryDNSACLService(opts...)
	r.Peers = NewSecondaryDNSPeerService(opts...)
	r.TSIGs = NewSecondaryDNSTSIGService(opts...)
	return
}
