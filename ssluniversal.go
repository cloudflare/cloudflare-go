// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SSLUniversalService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSSLUniversalService] method
// instead.
type SSLUniversalService struct {
	Options  []option.RequestOption
	Settings *SSLUniversalSettingService
}

// NewSSLUniversalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSSLUniversalService(opts ...option.RequestOption) (r *SSLUniversalService) {
	r = &SSLUniversalService{}
	r.Options = opts
	r.Settings = NewSSLUniversalSettingService(opts...)
	return
}
