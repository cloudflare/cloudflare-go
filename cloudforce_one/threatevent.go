// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ThreatEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventService] method instead.
type ThreatEventService struct {
	Options          []option.RequestOption
	Attackers        *ThreatEventAttackerService
	Categories       *ThreatEventCategoryService
	Countries        *ThreatEventCountryService
	Cron             *ThreatEventCronService
	Dataset          *ThreatEventDatasetService
	IndicatorTypes   *ThreatEventIndicatorTypeService
	Raw              *ThreatEventRawService
	Relate           *ThreatEventRelateService
	Tags             *ThreatEventTagService
	TargetIndustries *ThreatEventTargetIndustryService
	Insight          *ThreatEventInsightService
}

// NewThreatEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventService(opts ...option.RequestOption) (r *ThreatEventService) {
	r = &ThreatEventService{}
	r.Options = opts
	r.Attackers = NewThreatEventAttackerService(opts...)
	r.Categories = NewThreatEventCategoryService(opts...)
	r.Countries = NewThreatEventCountryService(opts...)
	r.Cron = NewThreatEventCronService(opts...)
	r.Dataset = NewThreatEventDatasetService(opts...)
	r.IndicatorTypes = NewThreatEventIndicatorTypeService(opts...)
	r.Raw = NewThreatEventRawService(opts...)
	r.Relate = NewThreatEventRelateService(opts...)
	r.Tags = NewThreatEventTagService(opts...)
	r.TargetIndustries = NewThreatEventTargetIndustryService(opts...)
	r.Insight = NewThreatEventInsightService(opts...)
	return
}
