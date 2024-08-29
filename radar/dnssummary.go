// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DNSSummaryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSSummaryService] method instead.
type DNSSummaryService struct {
	Options []option.RequestOption
}

// NewDNSSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSSummaryService(opts ...option.RequestOption) (r *DNSSummaryService) {
	r = &DNSSummaryService{}
	r.Options = opts
	return
}
