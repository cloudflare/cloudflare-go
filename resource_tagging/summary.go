// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_tagging

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// SummaryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSummaryService] method instead.
type SummaryService struct {
	Options []option.RequestOption
}

// NewSummaryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSummaryService(opts ...option.RequestOption) (r *SummaryService) {
	r = &SummaryService{}
	r.Options = opts
	return
}
