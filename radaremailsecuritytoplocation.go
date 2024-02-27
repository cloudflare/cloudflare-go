// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityTopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationService] method instead.
type RadarEmailSecurityTopLocationService struct {
	Options   []option.RequestOption
	ARC       *RadarEmailSecurityTopLocationARCService
	DKIM      *RadarEmailSecurityTopLocationDKIMService
	DMARC     *RadarEmailSecurityTopLocationDMARCService
	Malicious *RadarEmailSecurityTopLocationMaliciousService
	Spam      *RadarEmailSecurityTopLocationSpamService
	SPF       *RadarEmailSecurityTopLocationSPFService
}

// NewRadarEmailSecurityTopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationService) {
	r = &RadarEmailSecurityTopLocationService{}
	r.Options = opts
	r.ARC = NewRadarEmailSecurityTopLocationARCService(opts...)
	r.DKIM = NewRadarEmailSecurityTopLocationDKIMService(opts...)
	r.DMARC = NewRadarEmailSecurityTopLocationDMARCService(opts...)
	r.Malicious = NewRadarEmailSecurityTopLocationMaliciousService(opts...)
	r.Spam = NewRadarEmailSecurityTopLocationSpamService(opts...)
	r.SPF = NewRadarEmailSecurityTopLocationSPFService(opts...)
	return
}
