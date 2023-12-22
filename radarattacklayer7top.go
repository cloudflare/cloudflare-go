// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer7TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer7TopService]
// method instead.
type RadarAttackLayer7TopService struct {
	Options   []option.RequestOption
	Ases      *RadarAttackLayer7TopAseService
	Attacks   *RadarAttackLayer7TopAttackService
	Locations *RadarAttackLayer7TopLocationService
}

// NewRadarAttackLayer7TopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopService(opts ...option.RequestOption) (r *RadarAttackLayer7TopService) {
	r = &RadarAttackLayer7TopService{}
	r.Options = opts
	r.Ases = NewRadarAttackLayer7TopAseService(opts...)
	r.Attacks = NewRadarAttackLayer7TopAttackService(opts...)
	r.Locations = NewRadarAttackLayer7TopLocationService(opts...)
	return
}
