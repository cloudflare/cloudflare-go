// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer3TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer3TopService]
// method instead.
type RadarAttackLayer3TopService struct {
	Options   []option.RequestOption
	Attacks   *RadarAttackLayer3TopAttackService
	Industry  *RadarAttackLayer3TopIndustryService
	Locations *RadarAttackLayer3TopLocationService
	Vertical  *RadarAttackLayer3TopVerticalService
}

// NewRadarAttackLayer3TopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopService(opts ...option.RequestOption) (r *RadarAttackLayer3TopService) {
	r = &RadarAttackLayer3TopService{}
	r.Options = opts
	r.Attacks = NewRadarAttackLayer3TopAttackService(opts...)
	r.Industry = NewRadarAttackLayer3TopIndustryService(opts...)
	r.Locations = NewRadarAttackLayer3TopLocationService(opts...)
	r.Vertical = NewRadarAttackLayer3TopVerticalService(opts...)
	return
}
