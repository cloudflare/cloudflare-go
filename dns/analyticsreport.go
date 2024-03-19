// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnalyticsReportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsReportService] method
// instead.
type AnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *AnalyticsReportBytimeService
}

// NewAnalyticsReportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsReportService(opts ...option.RequestOption) (r *AnalyticsReportService) {
	r = &AnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewAnalyticsReportBytimeService(opts...)
	return
}
