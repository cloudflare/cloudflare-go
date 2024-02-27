// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityTopLocationMaliciousService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityTopLocationMaliciousService] method instead.
type RadarEmailSecurityTopLocationMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationMaliciousService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationMaliciousService) {
	r = &RadarEmailSecurityTopLocationMaliciousService{}
	r.Options = opts
	return
}
