// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarSearchService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarSearchService] method
// instead.
type RadarSearchService struct {
	Options []option.RequestOption
	Globals *RadarSearchGlobalService
}

// NewRadarSearchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarSearchService(opts ...option.RequestOption) (r *RadarSearchService) {
	r = &RadarSearchService{}
	r.Options = opts
	r.Globals = NewRadarSearchGlobalService(opts...)
	return
}
