// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarQualitySpeedService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualitySpeedService] method
// instead.
type RadarQualitySpeedService struct {
	Options   []option.RequestOption
	Histogram *RadarQualitySpeedHistogramService
	Summary   *RadarQualitySpeedSummaryService
	Top       *RadarQualitySpeedTopService
}

// NewRadarQualitySpeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedService(opts ...option.RequestOption) (r *RadarQualitySpeedService) {
	r = &RadarQualitySpeedService{}
	r.Options = opts
	r.Histogram = NewRadarQualitySpeedHistogramService(opts...)
	r.Summary = NewRadarQualitySpeedSummaryService(opts...)
	r.Top = NewRadarQualitySpeedTopService(opts...)
	return
}
