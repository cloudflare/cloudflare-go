// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DEXFleetStatusService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDEXFleetStatusService] method
// instead.
type DEXFleetStatusService struct {
	Options  []option.RequestOption
	Devices  *DEXFleetStatusDeviceService
	Live     *DEXFleetStatusLiveService
	OverTime *DEXFleetStatusOverTimeService
}

// NewDEXFleetStatusService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXFleetStatusService(opts ...option.RequestOption) (r *DEXFleetStatusService) {
	r = &DEXFleetStatusService{}
	r.Options = opts
	r.Devices = NewDEXFleetStatusDeviceService(opts...)
	r.Live = NewDEXFleetStatusLiveService(opts...)
	r.OverTime = NewDEXFleetStatusOverTimeService(opts...)
	return
}
