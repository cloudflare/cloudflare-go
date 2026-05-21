// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// DLSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLSService] method instead.
type DLSService struct {
	Options          []option.RequestOption
	Regions          *RegionService
	RegionalServices *RegionalServiceService
}

// NewDLSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDLSService(opts ...option.RequestOption) (r *DLSService) {
	r = &DLSService{}
	r.Options = opts
	r.Regions = NewRegionService(opts...)
	r.RegionalServices = NewRegionalServiceService(opts...)
	return
}
