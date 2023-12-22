// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarNetflowTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarNetflowTopService] method
// instead.
type RadarNetflowTopService struct {
	Options   []option.RequestOption
	Ases      *RadarNetflowTopAseService
	Locations *RadarNetflowTopLocationService
}

// NewRadarNetflowTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarNetflowTopService(opts ...option.RequestOption) (r *RadarNetflowTopService) {
	r = &RadarNetflowTopService{}
	r.Options = opts
	r.Ases = NewRadarNetflowTopAseService(opts...)
	r.Locations = NewRadarNetflowTopLocationService(opts...)
	return
}
