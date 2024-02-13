// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecuritySummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummaryService] method instead.
type RadarEmailSecuritySummaryService struct {
	Options          []option.RequestOption
	Arcs             *RadarEmailSecuritySummaryArcService
	DKIMs            *RadarEmailSecuritySummaryDKIMService
	Dmarcs           *RadarEmailSecuritySummaryDmarcService
	Malicious        *RadarEmailSecuritySummaryMaliciousService
	Spams            *RadarEmailSecuritySummarySpamService
	SPFs             *RadarEmailSecuritySummarySPFService
	ThreatCategories *RadarEmailSecuritySummaryThreatCategoryService
}

// NewRadarEmailSecuritySummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummaryService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryService) {
	r = &RadarEmailSecuritySummaryService{}
	r.Options = opts
	r.Arcs = NewRadarEmailSecuritySummaryArcService(opts...)
	r.DKIMs = NewRadarEmailSecuritySummaryDKIMService(opts...)
	r.Dmarcs = NewRadarEmailSecuritySummaryDmarcService(opts...)
	r.Malicious = NewRadarEmailSecuritySummaryMaliciousService(opts...)
	r.Spams = NewRadarEmailSecuritySummarySpamService(opts...)
	r.SPFs = NewRadarEmailSecuritySummarySPFService(opts...)
	r.ThreatCategories = NewRadarEmailSecuritySummaryThreatCategoryService(opts...)
	return
}
