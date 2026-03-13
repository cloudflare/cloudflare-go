// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// V2QueryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2QueryService] method instead.
type V2QueryService struct {
	Options []option.RequestOption
}

// NewV2QueryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2QueryService(opts ...option.RequestOption) (r *V2QueryService) {
	r = &V2QueryService{}
	r.Options = opts
	return
}
