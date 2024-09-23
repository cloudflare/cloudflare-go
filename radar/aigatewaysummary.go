// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// AIGatewaySummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewaySummaryService] method instead.
type AIGatewaySummaryService struct {
	Options []option.RequestOption
}

// NewAIGatewaySummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIGatewaySummaryService(opts ...option.RequestOption) (r *AIGatewaySummaryService) {
	r = &AIGatewaySummaryService{}
	r.Options = opts
	return
}
