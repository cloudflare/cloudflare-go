// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSecondaryDNSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSecondaryDNSService] method
// instead.
type ZoneSecondaryDNSService struct {
	Options    []option.RequestOption
	ForceAxfrs *ZoneSecondaryDNSForceAxfrService
	Incomings  *ZoneSecondaryDNSIncomingService
	Outgoings  *ZoneSecondaryDNSOutgoingService
}

// NewZoneSecondaryDNSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSecondaryDNSService(opts ...option.RequestOption) (r *ZoneSecondaryDNSService) {
	r = &ZoneSecondaryDNSService{}
	r.Options = opts
	r.ForceAxfrs = NewZoneSecondaryDNSForceAxfrService(opts...)
	r.Incomings = NewZoneSecondaryDNSIncomingService(opts...)
	r.Outgoings = NewZoneSecondaryDNSOutgoingService(opts...)
	return
}
