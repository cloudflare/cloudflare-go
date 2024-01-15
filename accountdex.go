// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDexService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountDexService] method instead.
type AccountDexService struct {
	Options               []option.RequestOption
	Colos                 *AccountDexColoService
	FleetStatusDevices    *AccountDexFleetStatusDeviceService
	FleetStatusLive       *AccountDexFleetStatusLiveService
	FleetStatusOverTime   *AccountDexFleetStatusOverTimeService
	HTTPTests             *AccountDexHTTPTestService
	Tests                 *AccountDexTestService
	TracerouteTestResults *AccountDexTracerouteTestResultService
	TracerouteTests       *AccountDexTracerouteTestService
}

// NewAccountDexService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDexService(opts ...option.RequestOption) (r *AccountDexService) {
	r = &AccountDexService{}
	r.Options = opts
	r.Colos = NewAccountDexColoService(opts...)
	r.FleetStatusDevices = NewAccountDexFleetStatusDeviceService(opts...)
	r.FleetStatusLive = NewAccountDexFleetStatusLiveService(opts...)
	r.FleetStatusOverTime = NewAccountDexFleetStatusOverTimeService(opts...)
	r.HTTPTests = NewAccountDexHTTPTestService(opts...)
	r.Tests = NewAccountDexTestService(opts...)
	r.TracerouteTestResults = NewAccountDexTracerouteTestResultService(opts...)
	r.TracerouteTests = NewAccountDexTracerouteTestService(opts...)
	return
}
