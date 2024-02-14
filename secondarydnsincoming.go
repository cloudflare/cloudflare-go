// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSIncomingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSIncomingService]
// method instead.
type SecondaryDNSIncomingService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSIncomingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSIncomingService(opts ...option.RequestOption) (r *SecondaryDNSIncomingService) {
	r = &SecondaryDNSIncomingService{}
	r.Options = opts
	return
}
