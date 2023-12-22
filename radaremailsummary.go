// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSummaryService] method
// instead.
type RadarEmailSummaryService struct {
	Options          []option.RequestOption
	Arcs             *RadarEmailSummaryArcService
	Dkims            *RadarEmailSummaryDkimService
	Dmarcs           *RadarEmailSummaryDmarcService
	Malicious        *RadarEmailSummaryMaliciousService
	Spams            *RadarEmailSummarySpamService
	Spfs             *RadarEmailSummarySpfService
	ThreatCategories *RadarEmailSummaryThreatCategoryService
}

// NewRadarEmailSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSummaryService(opts ...option.RequestOption) (r *RadarEmailSummaryService) {
	r = &RadarEmailSummaryService{}
	r.Options = opts
	r.Arcs = NewRadarEmailSummaryArcService(opts...)
	r.Dkims = NewRadarEmailSummaryDkimService(opts...)
	r.Dmarcs = NewRadarEmailSummaryDmarcService(opts...)
	r.Malicious = NewRadarEmailSummaryMaliciousService(opts...)
	r.Spams = NewRadarEmailSummarySpamService(opts...)
	r.Spfs = NewRadarEmailSummarySpfService(opts...)
	r.ThreatCategories = NewRadarEmailSummaryThreatCategoryService(opts...)
	return
}
