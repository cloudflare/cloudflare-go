// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer7TopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopLocationService] method instead.
type RadarAttackLayer7TopLocationService struct {
	Options []option.RequestOption
	Origins *RadarAttackLayer7TopLocationOriginService
	Targets *RadarAttackLayer7TopLocationTargetService
}

// NewRadarAttackLayer7TopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopLocationService(opts ...option.RequestOption) (r *RadarAttackLayer7TopLocationService) {
	r = &RadarAttackLayer7TopLocationService{}
	r.Options = opts
	r.Origins = NewRadarAttackLayer7TopLocationOriginService(opts...)
	r.Targets = NewRadarAttackLayer7TopLocationTargetService(opts...)
	return
}
