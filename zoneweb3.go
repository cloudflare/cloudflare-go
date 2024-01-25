// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWeb3Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneWeb3Service] method instead.
type ZoneWeb3Service struct {
	Options   []option.RequestOption
	Hostnames *ZoneWeb3HostnameService
}

// NewZoneWeb3Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneWeb3Service(opts ...option.RequestOption) (r *ZoneWeb3Service) {
	r = &ZoneWeb3Service{}
	r.Options = opts
	r.Hostnames = NewZoneWeb3HostnameService(opts...)
	return
}
