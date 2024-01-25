// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSecurityService] method
// instead.
type RadarEmailSecurityService struct {
	Options               []option.RequestOption
	ArcSummary            *RadarEmailSecurityArcSummaryService
	DkimSummary           *RadarEmailSecurityDkimSummaryService
	DmarcSummary          *RadarEmailSecurityDmarcSummaryService
	MaliciousSummary      *RadarEmailSecurityMaliciousSummaryService
	SpamSummary           *RadarEmailSecuritySpamSummaryService
	SpfSummary            *RadarEmailSecuritySpfSummaryService
	ThreatCategorySummary *RadarEmailSecurityThreatCategorySummaryService
	ArcTimeseries         *RadarEmailSecurityArcTimeseryService
	DkimTimeseries        *RadarEmailSecurityDkimTimeseryService
	DmarcTimeseries       *RadarEmailSecurityDmarcTimeseryService
	MaliciousTimeseries   *RadarEmailSecurityMaliciousTimeseryService
	SpamTimeseries        *RadarEmailSecuritySpamTimeseryService
	SpfTimeseries         *RadarEmailSecuritySpfTimeseryService
	TmeseriesGroups       *RadarEmailSecurityTmeseriesGroupService
	Top                   *RadarEmailSecurityTopService
}

// NewRadarEmailSecurityService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityService(opts ...option.RequestOption) (r *RadarEmailSecurityService) {
	r = &RadarEmailSecurityService{}
	r.Options = opts
	r.ArcSummary = NewRadarEmailSecurityArcSummaryService(opts...)
	r.DkimSummary = NewRadarEmailSecurityDkimSummaryService(opts...)
	r.DmarcSummary = NewRadarEmailSecurityDmarcSummaryService(opts...)
	r.MaliciousSummary = NewRadarEmailSecurityMaliciousSummaryService(opts...)
	r.SpamSummary = NewRadarEmailSecuritySpamSummaryService(opts...)
	r.SpfSummary = NewRadarEmailSecuritySpfSummaryService(opts...)
	r.ThreatCategorySummary = NewRadarEmailSecurityThreatCategorySummaryService(opts...)
	r.ArcTimeseries = NewRadarEmailSecurityArcTimeseryService(opts...)
	r.DkimTimeseries = NewRadarEmailSecurityDkimTimeseryService(opts...)
	r.DmarcTimeseries = NewRadarEmailSecurityDmarcTimeseryService(opts...)
	r.MaliciousTimeseries = NewRadarEmailSecurityMaliciousTimeseryService(opts...)
	r.SpamTimeseries = NewRadarEmailSecuritySpamTimeseryService(opts...)
	r.SpfTimeseries = NewRadarEmailSecuritySpfTimeseryService(opts...)
	r.TmeseriesGroups = NewRadarEmailSecurityTmeseriesGroupService(opts...)
	r.Top = NewRadarEmailSecurityTopService(opts...)
	return
}
