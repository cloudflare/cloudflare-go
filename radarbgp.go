// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarBgpService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarBgpService] method instead.
type RadarBgpService struct {
	Options    []option.RequestOption
	Leaks      *RadarBgpLeakService
	Timeseries *RadarBgpTimeseryService
	Tops       *RadarBgpTopService
}

// NewRadarBgpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpService(opts ...option.RequestOption) (r *RadarBgpService) {
	r = &RadarBgpService{}
	r.Options = opts
	r.Leaks = NewRadarBgpLeakService(opts...)
	r.Timeseries = NewRadarBgpTimeseryService(opts...)
	r.Tops = NewRadarBgpTopService(opts...)
	return
}
