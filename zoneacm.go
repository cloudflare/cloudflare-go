// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAcmService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneAcmService] method instead.
type ZoneAcmService struct {
	Options  []option.RequestOption
	TotalTls *ZoneAcmTotalTlService
}

// NewZoneAcmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneAcmService(opts ...option.RequestOption) (r *ZoneAcmService) {
	r = &ZoneAcmService{}
	r.Options = opts
	r.TotalTls = NewZoneAcmTotalTlService(opts...)
	return
}
