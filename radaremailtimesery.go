// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailTimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTimeseryService] method
// instead.
type RadarEmailTimeseryService struct {
	Options          []option.RequestOption
	Arcs             *RadarEmailTimeseryArcService
	Dkims            *RadarEmailTimeseryDkimService
	Dmarcs           *RadarEmailTimeseryDmarcService
	Malicious        *RadarEmailTimeseryMaliciousService
	Spams            *RadarEmailTimeserySpamService
	Spfs             *RadarEmailTimeserySpfService
	ThreatCategories *RadarEmailTimeseryThreatCategoryService
}

// NewRadarEmailTimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTimeseryService(opts ...option.RequestOption) (r *RadarEmailTimeseryService) {
	r = &RadarEmailTimeseryService{}
	r.Options = opts
	r.Arcs = NewRadarEmailTimeseryArcService(opts...)
	r.Dkims = NewRadarEmailTimeseryDkimService(opts...)
	r.Dmarcs = NewRadarEmailTimeseryDmarcService(opts...)
	r.Malicious = NewRadarEmailTimeseryMaliciousService(opts...)
	r.Spams = NewRadarEmailTimeserySpamService(opts...)
	r.Spfs = NewRadarEmailTimeserySpfService(opts...)
	r.ThreatCategories = NewRadarEmailTimeseryThreatCategoryService(opts...)
	return
}
