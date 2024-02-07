// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer3TopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TopLocationService] method instead.
type RadarAttackLayer3TopLocationService struct {
	Options []option.RequestOption
	Origin  *RadarAttackLayer3TopLocationOriginService
	Target  *RadarAttackLayer3TopLocationTargetService
}

// NewRadarAttackLayer3TopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopLocationService(opts ...option.RequestOption) (r *RadarAttackLayer3TopLocationService) {
	r = &RadarAttackLayer3TopLocationService{}
	r.Options = opts
	r.Origin = NewRadarAttackLayer3TopLocationOriginService(opts...)
	r.Target = NewRadarAttackLayer3TopLocationTargetService(opts...)
	return
}
