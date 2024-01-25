// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarRankingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarRankingService] method
// instead.
type RadarRankingService struct {
	Options          []option.RequestOption
	Domains          *RadarRankingDomainService
	TimeseriesGroups *RadarRankingTimeseriesGroupService
	Tops             *RadarRankingTopService
}

// NewRadarRankingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarRankingService(opts ...option.RequestOption) (r *RadarRankingService) {
	r = &RadarRankingService{}
	r.Options = opts
	r.Domains = NewRadarRankingDomainService(opts...)
	r.TimeseriesGroups = NewRadarRankingTimeseriesGroupService(opts...)
	r.Tops = NewRadarRankingTopService(opts...)
	return
}
