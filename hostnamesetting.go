// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// HostnameSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewHostnameSettingService] method
// instead.
type HostnameSettingService struct {
	Options []option.RequestOption
	TLS     *HostnameSettingTLSService
}

// NewHostnameSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHostnameSettingService(opts ...option.RequestOption) (r *HostnameSettingService) {
	r = &HostnameSettingService{}
	r.Options = opts
	r.TLS = NewHostnameSettingTLSService(opts...)
	return
}
