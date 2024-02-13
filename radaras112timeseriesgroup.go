// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112TimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupService] method instead.
type RadarAs112TimeseriesGroupService struct {
	Options   []option.RequestOption
	DNSSEC    *RadarAs112TimeseriesGroupDNSSECService
	Edns      *RadarAs112TimeseriesGroupEdnService
	IPVersion *RadarAs112TimeseriesGroupIPVersionService
}

// NewRadarAs112TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupService) {
	r = &RadarAs112TimeseriesGroupService{}
	r.Options = opts
	r.DNSSEC = NewRadarAs112TimeseriesGroupDNSSECService(opts...)
	r.Edns = NewRadarAs112TimeseriesGroupEdnService(opts...)
	r.IPVersion = NewRadarAs112TimeseriesGroupIPVersionService(opts...)
	return
}
