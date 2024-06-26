// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingTLSClientAuthService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingTLSClientAuthService] method instead.
type SettingTLSClientAuthService struct {
	Options []option.RequestOption
}

// NewSettingTLSClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingTLSClientAuthService(opts ...option.RequestOption) (r *SettingTLSClientAuthService) {
	r = &SettingTLSClientAuthService{}
	r.Options = opts
	return
}
