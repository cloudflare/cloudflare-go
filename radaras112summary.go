// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarAs112SummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112SummaryService] method
// instead.
type RadarAs112SummaryService struct {
	Options       []option.RequestOption
	Dnssecs       *RadarAs112SummaryDnssecService
	Edns          *RadarAs112SummaryEdnService
	IPVersions    *RadarAs112SummaryIPVersionService
	Protocols     *RadarAs112SummaryProtocolService
	QueryTypes    *RadarAs112SummaryQueryTypeService
	ResponseCodes *RadarAs112SummaryResponseCodeService
}

// NewRadarAs112SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryService(opts ...option.RequestOption) (r *RadarAs112SummaryService) {
	r = &RadarAs112SummaryService{}
	r.Options = opts
	r.Dnssecs = NewRadarAs112SummaryDnssecService(opts...)
	r.Edns = NewRadarAs112SummaryEdnService(opts...)
	r.IPVersions = NewRadarAs112SummaryIPVersionService(opts...)
	r.Protocols = NewRadarAs112SummaryProtocolService(opts...)
	r.QueryTypes = NewRadarAs112SummaryQueryTypeService(opts...)
	r.ResponseCodes = NewRadarAs112SummaryResponseCodeService(opts...)
	return
}
