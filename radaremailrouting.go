// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarEmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailRoutingService] method
// instead.
type RadarEmailRoutingService struct {
	Options          []option.RequestOption
	Summary          *RadarEmailRoutingSummaryService
	TimeseriesGroups *RadarEmailRoutingTimeseriesGroupService
}

// NewRadarEmailRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailRoutingService(opts ...option.RequestOption) (r *RadarEmailRoutingService) {
	r = &RadarEmailRoutingService{}
	r.Options = opts
	r.Summary = NewRadarEmailRoutingSummaryService(opts...)
	r.TimeseriesGroups = NewRadarEmailRoutingTimeseriesGroupService(opts...)
	return
}
