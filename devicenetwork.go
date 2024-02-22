// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DeviceNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceNetworkService] method
// instead.
type DeviceNetworkService struct {
	Options []option.RequestOption
}

// NewDeviceNetworkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceNetworkService(opts ...option.RequestOption) (r *DeviceNetworkService) {
	r = &DeviceNetworkService{}
	r.Options = opts
	return
}
