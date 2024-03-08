// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// DNSFirewallAnalyticsService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDNSFirewallAnalyticsService]
// method instead.
type DNSFirewallAnalyticsService struct {
	Options []option.RequestOption
	Reports *DNSFirewallAnalyticsReportService
}

// NewDNSFirewallAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDNSFirewallAnalyticsService(opts ...option.RequestOption) (r *DNSFirewallAnalyticsService) {
	r = &DNSFirewallAnalyticsService{}
	r.Options = opts
	r.Reports = NewDNSFirewallAnalyticsReportService(opts...)
	return
}
