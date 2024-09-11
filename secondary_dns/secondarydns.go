// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secondary_dns

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SecondaryDNSService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecondaryDNSService] method instead.
type SecondaryDNSService struct {
	Options   []option.RequestOption
	ForceAXFR *ForceAXFRService
	Incoming  *IncomingService
	Outgoing  *OutgoingService
	ACLs      *ACLService
	Peers     *PeerService
	TSIGs     *TSIGService
}

// NewSecondaryDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecondaryDNSService(opts ...option.RequestOption) (r *SecondaryDNSService) {
	r = &SecondaryDNSService{}
	r.Options = opts
	r.ForceAXFR = NewForceAXFRService(opts...)
	r.Incoming = NewIncomingService(opts...)
	r.Outgoing = NewOutgoingService(opts...)
	r.ACLs = NewACLService(opts...)
	r.Peers = NewPeerService(opts...)
	r.TSIGs = NewTSIGService(opts...)
	return
}
