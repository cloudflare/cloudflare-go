// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DevicePolicyExcludeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDevicePolicyExcludeService]
// method instead.
type DevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewDevicePolicyExcludeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevicePolicyExcludeService(opts ...option.RequestOption) (r *DevicePolicyExcludeService) {
	r = &DevicePolicyExcludeService{}
	r.Options = opts
	return
}
