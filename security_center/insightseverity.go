// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InsightSeverityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightSeverityService] method instead.
type InsightSeverityService struct {
	Options []option.RequestOption
}

// NewInsightSeverityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightSeverityService(opts ...option.RequestOption) (r *InsightSeverityService) {
	r = &InsightSeverityService{}
	r.Options = opts
	return
}
