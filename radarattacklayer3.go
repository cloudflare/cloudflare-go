// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer3Service] method
// instead.
type RadarAttackLayer3Service struct {
	Options          []option.RequestOption
	TimeseriesGroups *RadarAttackLayer3TimeseriesGroupService
	Top              *RadarAttackLayer3TopService
}

// NewRadarAttackLayer3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3Service(opts ...option.RequestOption) (r *RadarAttackLayer3Service) {
	r = &RadarAttackLayer3Service{}
	r.Options = opts
	r.TimeseriesGroups = NewRadarAttackLayer3TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer3TopService(opts...)
	return
}
