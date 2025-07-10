// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// MatchService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatchService] method instead.
type MatchService struct {
	Options []option.RequestOption
}

// NewMatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMatchService(opts ...option.RequestOption) (r *MatchService) {
	r = &MatchService{}
	r.Options = opts
	return
}
