// File generated from our OpenAPI spec by Stainless.

package hostnames

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSettingService] method instead.
type SettingService struct {
	Options []option.RequestOption
	TLS     *SettingTLSService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	r.TLS = NewSettingTLSService(opts...)
	return
}
