// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarBgpRouteService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpRouteService] method
// instead.
type RadarBgpRouteService struct {
	Options []option.RequestOption
	Moas    *RadarBgpRouteMoaService
	Pfx2as  *RadarBgpRoutePfx2aService
	Stats   *RadarBgpRouteStatService
}

// NewRadarBgpRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpRouteService(opts ...option.RequestOption) (r *RadarBgpRouteService) {
	r = &RadarBgpRouteService{}
	r.Options = opts
	r.Moas = NewRadarBgpRouteMoaService(opts...)
	r.Pfx2as = NewRadarBgpRoutePfx2aService(opts...)
	r.Stats = NewRadarBgpRouteStatService(opts...)
	return
}
