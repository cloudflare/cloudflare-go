// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneArgoService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneArgoService] method instead.
type ZoneArgoService struct {
	Options       []option.RequestOption
	SmartRoutings *ZoneArgoSmartRoutingService
	TieredCaching *ZoneArgoTieredCachingService
}

// NewZoneArgoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneArgoService(opts ...option.RequestOption) (r *ZoneArgoService) {
	r = &ZoneArgoService{}
	r.Options = opts
	r.SmartRoutings = NewZoneArgoSmartRoutingService(opts...)
	r.TieredCaching = NewZoneArgoTieredCachingService(opts...)
	return
}
