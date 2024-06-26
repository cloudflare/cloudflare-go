// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// SettingHTTP2Service contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingHTTP2Service] method instead.
type SettingHTTP2Service struct {
	Options []option.RequestOption
}

// NewSettingHTTP2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingHTTP2Service(opts ...option.RequestOption) (r *SettingHTTP2Service) {
	r = &SettingHTTP2Service{}
	r.Options = opts
	return
}
