// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPTimeseriesGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupService] method instead.
type RadarHTTPTimeseriesGroupService struct {
	Options         []option.RequestOption
	ByBotClass      *RadarHTTPTimeseriesGroupByBotClassService
	ByBrowser       *RadarHTTPTimeseriesGroupByBrowserService
	ByBrowserFamily *RadarHTTPTimeseriesGroupByBrowserFamilyService
	ByDeviceType    *RadarHTTPTimeseriesGroupByDeviceTypeService
	ByHTTPProtocol  *RadarHTTPTimeseriesGroupByHTTPProtocolService
	ByHTTPVersion   *RadarHTTPTimeseriesGroupByHTTPVersionService
	ByIPVersion     *RadarHTTPTimeseriesGroupByIPVersionService
	ByOs            *RadarHTTPTimeseriesGroupByOService
	ByTlsVersion    *RadarHTTPTimeseriesGroupByTlsVersionService
}

// NewRadarHTTPTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseriesGroupService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupService) {
	r = &RadarHTTPTimeseriesGroupService{}
	r.Options = opts
	r.ByBotClass = NewRadarHTTPTimeseriesGroupByBotClassService(opts...)
	r.ByBrowser = NewRadarHTTPTimeseriesGroupByBrowserService(opts...)
	r.ByBrowserFamily = NewRadarHTTPTimeseriesGroupByBrowserFamilyService(opts...)
	r.ByDeviceType = NewRadarHTTPTimeseriesGroupByDeviceTypeService(opts...)
	r.ByHTTPProtocol = NewRadarHTTPTimeseriesGroupByHTTPProtocolService(opts...)
	r.ByHTTPVersion = NewRadarHTTPTimeseriesGroupByHTTPVersionService(opts...)
	r.ByIPVersion = NewRadarHTTPTimeseriesGroupByIPVersionService(opts...)
	r.ByOs = NewRadarHTTPTimeseriesGroupByOService(opts...)
	r.ByTlsVersion = NewRadarHTTPTimeseriesGroupByTlsVersionService(opts...)
	return
}
