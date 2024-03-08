// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// SpectrumAnalyticsService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSpectrumAnalyticsService] method
// instead.
type SpectrumAnalyticsService struct {
	Options    []option.RequestOption
	Aggregates *SpectrumAnalyticsAggregateService
	Events     *SpectrumAnalyticsEventService
}

// NewSpectrumAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpectrumAnalyticsService(opts ...option.RequestOption) (r *SpectrumAnalyticsService) {
	r = &SpectrumAnalyticsService{}
	r.Options = opts
	r.Aggregates = NewSpectrumAnalyticsAggregateService(opts...)
	r.Events = NewSpectrumAnalyticsEventService(opts...)
	return
}
