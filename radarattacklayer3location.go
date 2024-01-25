// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer3LocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3LocationService] method instead.
type RadarAttackLayer3LocationService struct {
	Options []option.RequestOption
	Origin  *RadarAttackLayer3LocationOriginService
	Target  *RadarAttackLayer3LocationTargetService
}

// NewRadarAttackLayer3LocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3LocationService(opts ...option.RequestOption) (r *RadarAttackLayer3LocationService) {
	r = &RadarAttackLayer3LocationService{}
	r.Options = opts
	r.Origin = NewRadarAttackLayer3LocationOriginService(opts...)
	r.Target = NewRadarAttackLayer3LocationTargetService(opts...)
	return
}
