// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InsightService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightService] method instead.
type InsightService struct {
	Options  []option.RequestOption
	Class    *InsightClassService
	Severity *InsightSeverityService
	Type     *InsightTypeService
}

// NewInsightService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInsightService(opts ...option.RequestOption) (r *InsightService) {
	r = &InsightService{}
	r.Options = opts
	r.Class = NewInsightClassService(opts...)
	r.Severity = NewInsightSeverityService(opts...)
	r.Type = NewInsightTypeService(opts...)
	return
}
