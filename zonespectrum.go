// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSpectrumService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSpectrumService] method
// instead.
type ZoneSpectrumService struct {
	Options   []option.RequestOption
	Analytics *ZoneSpectrumAnalyticsService
}

// NewZoneSpectrumService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSpectrumService(opts ...option.RequestOption) (r *ZoneSpectrumService) {
	r = &ZoneSpectrumService{}
	r.Options = opts
	r.Analytics = NewZoneSpectrumAnalyticsService(opts...)
	return
}
