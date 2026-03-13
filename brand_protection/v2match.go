// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// V2MatchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV2MatchService] method instead.
type V2MatchService struct {
	Options []option.RequestOption
}

// NewV2MatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2MatchService(opts ...option.RequestOption) (r *V2MatchService) {
	r = &V2MatchService{}
	r.Options = opts
	return
}
