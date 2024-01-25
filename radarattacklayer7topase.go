// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer7TopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7TopAseService] method instead.
type RadarAttackLayer7TopAseService struct {
	Options []option.RequestOption
	Origins *RadarAttackLayer7TopAseOriginService
}

// NewRadarAttackLayer7TopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopAseService(opts ...option.RequestOption) (r *RadarAttackLayer7TopAseService) {
	r = &RadarAttackLayer7TopAseService{}
	r.Options = opts
	r.Origins = NewRadarAttackLayer7TopAseOriginService(opts...)
	return
}
