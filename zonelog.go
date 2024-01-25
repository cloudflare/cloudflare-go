// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneLogService] method instead.
type ZoneLogService struct {
	Options   []option.RequestOption
	Controls  *ZoneLogControlService
	Rayids    *ZoneLogRayidService
	Receiveds *ZoneLogReceivedService
}

// NewZoneLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneLogService(opts ...option.RequestOption) (r *ZoneLogService) {
	r = &ZoneLogService{}
	r.Options = opts
	r.Controls = NewZoneLogControlService(opts...)
	r.Rayids = NewZoneLogRayidService(opts...)
	r.Receiveds = NewZoneLogReceivedService(opts...)
	return
}
