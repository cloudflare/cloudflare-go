// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// AnalyticsService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options []option.RequestOption
	Reports *AnalyticsReportService
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r *AnalyticsService) {
	r = &AnalyticsService{}
	r.Options = opts
	r.Reports = NewAnalyticsReportService(opts...)
	return
}

// Unit of time to group data by.
type Delta string

const (
	DeltaAll        Delta = "all"
	DeltaAuto       Delta = "auto"
	DeltaYear       Delta = "year"
	DeltaQuarter    Delta = "quarter"
	DeltaMonth      Delta = "month"
	DeltaWeek       Delta = "week"
	DeltaDay        Delta = "day"
	DeltaHour       Delta = "hour"
	DeltaDekaminute Delta = "dekaminute"
	DeltaMinute     Delta = "minute"
)

func (r Delta) IsKnown() bool {
	switch r {
	case DeltaAll, DeltaAuto, DeltaYear, DeltaQuarter, DeltaMonth, DeltaWeek, DeltaDay, DeltaHour, DeltaDekaminute, DeltaMinute:
		return true
	}
	return false
}
