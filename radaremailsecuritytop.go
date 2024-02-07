// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailSecurityTopService]
// method instead.
type RadarEmailSecurityTopService struct {
	Options []option.RequestOption
	Ases    *RadarEmailSecurityTopAseService
}

// NewRadarEmailSecurityTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopService(opts ...option.RequestOption) (r *RadarEmailSecurityTopService) {
	r = &RadarEmailSecurityTopService{}
	r.Options = opts
	r.Ases = NewRadarEmailSecurityTopAseService(opts...)
	return
}
