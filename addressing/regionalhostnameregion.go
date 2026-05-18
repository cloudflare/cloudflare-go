// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// RegionalHostnameRegionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalHostnameRegionService] method instead.
type RegionalHostnameRegionService struct {
	Options []option.RequestOption
}

// NewRegionalHostnameRegionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegionalHostnameRegionService(opts ...option.RequestOption) (r *RegionalHostnameRegionService) {
	r = &RegionalHostnameRegionService{}
	r.Options = opts
	return
}
