// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarEmailService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarEmailService] method instead.
type RadarEmailService struct {
	Options  []option.RequestOption
	Routing  *RadarEmailRoutingService
	Security *RadarEmailSecurityService
}

// NewRadarEmailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEmailService(opts ...option.RequestOption) (r *RadarEmailService) {
	r = &RadarEmailService{}
	r.Options = opts
	r.Routing = NewRadarEmailRoutingService(opts...)
	r.Security = NewRadarEmailSecurityService(opts...)
	return
}
