// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112TopService] method
// instead.
type RadarAs112TopService struct {
	Options   []option.RequestOption
	Locations *RadarAs112TopLocationService
}

// NewRadarAs112TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAs112TopService(opts ...option.RequestOption) (r *RadarAs112TopService) {
	r = &RadarAs112TopService{}
	r.Options = opts
	r.Locations = NewRadarAs112TopLocationService(opts...)
	return
}
