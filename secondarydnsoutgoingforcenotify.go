// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSOutgoingForceNotifyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSecondaryDNSOutgoingForceNotifyService] method instead.
type SecondaryDNSOutgoingForceNotifyService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSOutgoingForceNotifyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingForceNotifyService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingForceNotifyService) {
	r = &SecondaryDNSOutgoingForceNotifyService{}
	r.Options = opts
	return
}
