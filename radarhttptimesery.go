// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPTimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTimeseryService] method
// instead.
type RadarHTTPTimeseryService struct {
	Options         []option.RequestOption
	BotClasses      *RadarHTTPTimeseryBotClassService
	Browsers        *RadarHTTPTimeseryBrowserService
	BrowserFamilies *RadarHTTPTimeseryBrowserFamilyService
	DeviceTypes     *RadarHTTPTimeseryDeviceTypeService
	HTTPProtocols   *RadarHTTPTimeseryHTTPProtocolService
	HTTPVersions    *RadarHTTPTimeseryHTTPVersionService
	IPVersions      *RadarHTTPTimeseryIPVersionService
	Os              *RadarHTTPTimeseryOService
	TlsVersions     *RadarHTTPTimeseryTlsVersionService
}

// NewRadarHTTPTimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryService(opts ...option.RequestOption) (r *RadarHTTPTimeseryService) {
	r = &RadarHTTPTimeseryService{}
	r.Options = opts
	r.BotClasses = NewRadarHTTPTimeseryBotClassService(opts...)
	r.Browsers = NewRadarHTTPTimeseryBrowserService(opts...)
	r.BrowserFamilies = NewRadarHTTPTimeseryBrowserFamilyService(opts...)
	r.DeviceTypes = NewRadarHTTPTimeseryDeviceTypeService(opts...)
	r.HTTPProtocols = NewRadarHTTPTimeseryHTTPProtocolService(opts...)
	r.HTTPVersions = NewRadarHTTPTimeseryHTTPVersionService(opts...)
	r.IPVersions = NewRadarHTTPTimeseryIPVersionService(opts...)
	r.Os = NewRadarHTTPTimeseryOService(opts...)
	r.TlsVersions = NewRadarHTTPTimeseryTlsVersionService(opts...)
	return
}
