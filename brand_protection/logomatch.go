// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// LogoMatchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogoMatchService] method instead.
type LogoMatchService struct {
	Options []option.RequestOption
}

// NewLogoMatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogoMatchService(opts ...option.RequestOption) (r *LogoMatchService) {
	r = &LogoMatchService{}
	r.Options = opts
	return
}
