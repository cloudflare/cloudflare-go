// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WhoisService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWhoisService] method instead.
type WhoisService struct {
	Options []option.RequestOption
}

// NewWhoisService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWhoisService(opts ...option.RequestOption) (r *WhoisService) {
	r = &WhoisService{}
	r.Options = opts
	return
}
