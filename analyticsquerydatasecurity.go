// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// AnalyticsQueryDataSecurityService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsQueryDataSecurityService] method instead.
type AnalyticsQueryDataSecurityService struct {
	Options         []option.RequestOption
	ContentFindings *AnalyticsQueryDataSecurityContentFindingService
	Findings        *AnalyticsQueryDataSecurityFindingService
}

// NewAnalyticsQueryDataSecurityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAnalyticsQueryDataSecurityService(opts ...option.RequestOption) (r *AnalyticsQueryDataSecurityService) {
	r = &AnalyticsQueryDataSecurityService{}
	r.Options = opts
	r.ContentFindings = NewAnalyticsQueryDataSecurityContentFindingService(opts...)
	r.Findings = NewAnalyticsQueryDataSecurityFindingService(opts...)
	return
}
