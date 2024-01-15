// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAttackLayer7Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer7Service] method
// instead.
type RadarAttackLayer7Service struct {
	Options            []option.RequestOption
	Summaries          *RadarAttackLayer7SummaryService
	MitigationProducts *RadarAttackLayer7MitigationProductService
	HTTPMethods        *RadarAttackLayer7HTTPMethodService
	HTTPVersions       *RadarAttackLayer7HTTPVersionService
	Industries         *RadarAttackLayer7IndustryService
	IPVersions         *RadarAttackLayer7IPVersionService
	ManagedRules       *RadarAttackLayer7ManagedRuleService
	Verticals          *RadarAttackLayer7VerticalService
	Timeseries         *RadarAttackLayer7TimeseryService
	TimeseriesGroups   *RadarAttackLayer7TimeseriesGroupService
	Tops               *RadarAttackLayer7TopService
}

// NewRadarAttackLayer7Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7Service(opts ...option.RequestOption) (r *RadarAttackLayer7Service) {
	r = &RadarAttackLayer7Service{}
	r.Options = opts
	r.Summaries = NewRadarAttackLayer7SummaryService(opts...)
	r.MitigationProducts = NewRadarAttackLayer7MitigationProductService(opts...)
	r.HTTPMethods = NewRadarAttackLayer7HTTPMethodService(opts...)
	r.HTTPVersions = NewRadarAttackLayer7HTTPVersionService(opts...)
	r.Industries = NewRadarAttackLayer7IndustryService(opts...)
	r.IPVersions = NewRadarAttackLayer7IPVersionService(opts...)
	r.ManagedRules = NewRadarAttackLayer7ManagedRuleService(opts...)
	r.Verticals = NewRadarAttackLayer7VerticalService(opts...)
	r.Timeseries = NewRadarAttackLayer7TimeseryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer7TimeseriesGroupService(opts...)
	r.Tops = NewRadarAttackLayer7TopService(opts...)
	return
}
