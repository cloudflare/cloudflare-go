// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityTopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseService] method instead.
type RadarEmailSecurityTopAseService struct {
	Options   []option.RequestOption
	ARC       *RadarEmailSecurityTopAseARCService
	DKIM      *RadarEmailSecurityTopAseDKIMService
	DMARC     *RadarEmailSecurityTopAseDMARCService
	Malicious *RadarEmailSecurityTopAseMaliciousService
	Spam      *RadarEmailSecurityTopAseSpamService
	SPF       *RadarEmailSecurityTopAseSPFService
}

// NewRadarEmailSecurityTopAseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseService) {
	r = &RadarEmailSecurityTopAseService{}
	r.Options = opts
	r.ARC = NewRadarEmailSecurityTopAseARCService(opts...)
	r.DKIM = NewRadarEmailSecurityTopAseDKIMService(opts...)
	r.DMARC = NewRadarEmailSecurityTopAseDMARCService(opts...)
	r.Malicious = NewRadarEmailSecurityTopAseMaliciousService(opts...)
	r.Spam = NewRadarEmailSecurityTopAseSpamService(opts...)
	r.SPF = NewRadarEmailSecurityTopAseSPFService(opts...)
	return
}
