// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// FirewallAnalyticsReportBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewFirewallAnalyticsReportBytimeService] method instead.
type FirewallAnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewFirewallAnalyticsReportBytimeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFirewallAnalyticsReportBytimeService(opts ...option.RequestOption) (r *FirewallAnalyticsReportBytimeService) {
	r = &FirewallAnalyticsReportBytimeService{}
	r.Options = opts
	return
}
