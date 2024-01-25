// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneWorkerService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneWorkerService] method instead.
type ZoneWorkerService struct {
	Options []option.RequestOption
	Scripts *ZoneWorkerScriptService
	Filters *ZoneWorkerFilterService
	Routes  *ZoneWorkerRouteService
}

// NewZoneWorkerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneWorkerService(opts ...option.RequestOption) (r *ZoneWorkerService) {
	r = &ZoneWorkerService{}
	r.Options = opts
	r.Scripts = NewZoneWorkerScriptService(opts...)
	r.Filters = NewZoneWorkerFilterService(opts...)
	r.Routes = NewZoneWorkerRouteService(opts...)
	return
}
