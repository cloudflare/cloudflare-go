// File generated from our OpenAPI spec by Stainless.

package dns

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// FirewallAnalyticsService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallAnalyticsService] method
// instead.
type FirewallAnalyticsService struct {
	Options []option.RequestOption
	Reports *FirewallAnalyticsReportService
}

// NewFirewallAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallAnalyticsService(opts ...option.RequestOption) (r *FirewallAnalyticsService) {
	r = &FirewallAnalyticsService{}
	r.Options = opts
	r.Reports = NewFirewallAnalyticsReportService(opts...)
	return
}
