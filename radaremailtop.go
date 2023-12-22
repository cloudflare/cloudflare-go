// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopService] method
// instead.
type RadarEmailTopService struct {
	Options   []option.RequestOption
	Ases      *RadarEmailTopAseService
	Locations *RadarEmailTopLocationService
}

// NewRadarEmailTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarEmailTopService(opts ...option.RequestOption) (r *RadarEmailTopService) {
	r = &RadarEmailTopService{}
	r.Options = opts
	r.Ases = NewRadarEmailTopAseService(opts...)
	r.Locations = NewRadarEmailTopLocationService(opts...)
	return
}
