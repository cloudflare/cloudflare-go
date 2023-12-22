// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPSummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPSummaryService] method
// instead.
type RadarHTTPSummaryService struct {
	Options       []option.RequestOption
	BotClasses    *RadarHTTPSummaryBotClassService
	DeviceTypes   *RadarHTTPSummaryDeviceTypeService
	HTTPProtocols *RadarHTTPSummaryHTTPProtocolService
	HTTPVersions  *RadarHTTPSummaryHTTPVersionService
	IPVersions    *RadarHTTPSummaryIPVersionService
	Os            *RadarHTTPSummaryOService
	TlsVersions   *RadarHTTPSummaryTlsVersionService
}

// NewRadarHTTPSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPSummaryService(opts ...option.RequestOption) (r *RadarHTTPSummaryService) {
	r = &RadarHTTPSummaryService{}
	r.Options = opts
	r.BotClasses = NewRadarHTTPSummaryBotClassService(opts...)
	r.DeviceTypes = NewRadarHTTPSummaryDeviceTypeService(opts...)
	r.HTTPProtocols = NewRadarHTTPSummaryHTTPProtocolService(opts...)
	r.HTTPVersions = NewRadarHTTPSummaryHTTPVersionService(opts...)
	r.IPVersions = NewRadarHTTPSummaryIPVersionService(opts...)
	r.Os = NewRadarHTTPSummaryOService(opts...)
	r.TlsVersions = NewRadarHTTPSummaryTlsVersionService(opts...)
	return
}
