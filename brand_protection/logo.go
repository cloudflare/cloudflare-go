// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// LogoService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogoService] method instead.
type LogoService struct {
	Options []option.RequestOption
}

// NewLogoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogoService(opts ...option.RequestOption) (r *LogoService) {
	r = &LogoService{}
	r.Options = opts
	return
}
