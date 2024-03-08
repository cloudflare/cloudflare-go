// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// DEXService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDEXService] method instead.
type DEXService struct {
	Options               []option.RequestOption
	Colos                 *DEXColoService
	FleetStatus           *DEXFleetStatusService
	HTTPTests             *DEXHTTPTestService
	Tests                 *DEXTestService
	TracerouteTestResults *DEXTracerouteTestResultService
	TracerouteTests       *DEXTracerouteTestService
}

// NewDEXService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXService(opts ...option.RequestOption) (r *DEXService) {
	r = &DEXService{}
	r.Options = opts
	r.Colos = NewDEXColoService(opts...)
	r.FleetStatus = NewDEXFleetStatusService(opts...)
	r.HTTPTests = NewDEXHTTPTestService(opts...)
	r.Tests = NewDEXTestService(opts...)
	r.TracerouteTestResults = NewDEXTracerouteTestResultService(opts...)
	r.TracerouteTests = NewDEXTracerouteTestService(opts...)
	return
}
