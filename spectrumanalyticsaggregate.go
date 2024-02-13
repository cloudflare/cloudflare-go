// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SpectrumAnalyticsAggregateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewSpectrumAnalyticsAggregateService] method instead.
type SpectrumAnalyticsAggregateService struct {
	Options  []option.RequestOption
	Currents *SpectrumAnalyticsAggregateCurrentService
}

// NewSpectrumAnalyticsAggregateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSpectrumAnalyticsAggregateService(opts ...option.RequestOption) (r *SpectrumAnalyticsAggregateService) {
	r = &SpectrumAnalyticsAggregateService{}
	r.Options = opts
	r.Currents = NewSpectrumAnalyticsAggregateCurrentService(opts...)
	return
}
