// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarEmailSecurityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSecurityService] method
// instead.
type RadarEmailSecurityService struct {
	Options          []option.RequestOption
	Top              *RadarEmailSecurityTopService
	Summary          *RadarEmailSecuritySummaryService
	TimeseriesGroups *RadarEmailSecurityTimeseriesGroupService
}

// NewRadarEmailSecurityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityService(opts ...option.RequestOption) (r *RadarEmailSecurityService) {
	r = &RadarEmailSecurityService{}
	r.Options = opts
	r.Top = NewRadarEmailSecurityTopService(opts...)
	r.Summary = NewRadarEmailSecuritySummaryService(opts...)
	r.TimeseriesGroups = NewRadarEmailSecurityTimeseriesGroupService(opts...)
	return
}
