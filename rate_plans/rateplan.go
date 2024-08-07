// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_plans

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RatePlanService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRatePlanService] method instead.
type RatePlanService struct {
	Options []option.RequestOption
}

// NewRatePlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatePlanService(opts ...option.RequestOption) (r *RatePlanService) {
	r = &RatePlanService{}
	r.Options = opts
	return
}
