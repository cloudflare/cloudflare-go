// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSOutgoingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSOutgoingService]
// method instead.
type SecondaryDNSOutgoingService struct {
	Options       []option.RequestOption
	Disables      *SecondaryDNSOutgoingDisableService
	Enables       *SecondaryDNSOutgoingEnableService
	ForceNotifies *SecondaryDNSOutgoingForceNotifyService
	Statuses      *SecondaryDNSOutgoingStatusService
}

// NewSecondaryDNSOutgoingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingService) {
	r = &SecondaryDNSOutgoingService{}
	r.Options = opts
	r.Disables = NewSecondaryDNSOutgoingDisableService(opts...)
	r.Enables = NewSecondaryDNSOutgoingEnableService(opts...)
	r.ForceNotifies = NewSecondaryDNSOutgoingForceNotifyService(opts...)
	r.Statuses = NewSecondaryDNSOutgoingStatusService(opts...)
	return
}
