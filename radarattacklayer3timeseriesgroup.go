// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer3TimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TimeseriesGroupService] method instead.
type RadarAttackLayer3TimeseriesGroupService struct {
	Options   []option.RequestOption
	Industry  *RadarAttackLayer3TimeseriesGroupIndustryService
	IPVersion *RadarAttackLayer3TimeseriesGroupIPVersionService
	Protocol  *RadarAttackLayer3TimeseriesGroupProtocolService
	Vector    *RadarAttackLayer3TimeseriesGroupVectorService
	Vertical  *RadarAttackLayer3TimeseriesGroupVerticalService
}

// NewRadarAttackLayer3TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupService) {
	r = &RadarAttackLayer3TimeseriesGroupService{}
	r.Options = opts
	r.Industry = NewRadarAttackLayer3TimeseriesGroupIndustryService(opts...)
	r.IPVersion = NewRadarAttackLayer3TimeseriesGroupIPVersionService(opts...)
	r.Protocol = NewRadarAttackLayer3TimeseriesGroupProtocolService(opts...)
	r.Vector = NewRadarAttackLayer3TimeseriesGroupVectorService(opts...)
	r.Vertical = NewRadarAttackLayer3TimeseriesGroupVerticalService(opts...)
	return
}
