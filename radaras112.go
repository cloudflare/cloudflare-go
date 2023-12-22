// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarAs112Service] method instead.
type RadarAs112Service struct {
	Options    []option.RequestOption
	Summaries  *RadarAs112SummaryService
	Timeseries *RadarAs112TimeseryService
	Tops       *RadarAs112TopService
}

// NewRadarAs112Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAs112Service(opts ...option.RequestOption) (r *RadarAs112Service) {
	r = &RadarAs112Service{}
	r.Options = opts
	r.Summaries = NewRadarAs112SummaryService(opts...)
	r.Timeseries = NewRadarAs112TimeseryService(opts...)
	r.Tops = NewRadarAs112TopService(opts...)
	return
}
