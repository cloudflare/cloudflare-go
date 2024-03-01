// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDEXFleetStatusService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustDEXFleetStatusService] method instead.
type ZeroTrustDEXFleetStatusService struct {
	Options  []option.RequestOption
	Devices  *ZeroTrustDEXFleetStatusDeviceService
	Live     *ZeroTrustDEXFleetStatusLiveService
	OverTime *ZeroTrustDEXFleetStatusOverTimeService
}

// NewZeroTrustDEXFleetStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDEXFleetStatusService(opts ...option.RequestOption) (r *ZeroTrustDEXFleetStatusService) {
	r = &ZeroTrustDEXFleetStatusService{}
	r.Options = opts
	r.Devices = NewZeroTrustDEXFleetStatusDeviceService(opts...)
	r.Live = NewZeroTrustDEXFleetStatusLiveService(opts...)
	r.OverTime = NewZeroTrustDEXFleetStatusOverTimeService(opts...)
	return
}
