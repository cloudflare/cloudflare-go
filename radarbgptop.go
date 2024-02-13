// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarBGPTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBGPTopService] method
// instead.
type RadarBGPTopService struct {
	Options  []option.RequestOption
	Ases     *RadarBGPTopAseService
	Prefixes *RadarBGPTopPrefixService
}

// NewRadarBGPTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBGPTopService(opts ...option.RequestOption) (r *RadarBGPTopService) {
	r = &RadarBGPTopService{}
	r.Options = opts
	r.Ases = NewRadarBGPTopAseService(opts...)
	r.Prefixes = NewRadarBGPTopPrefixService(opts...)
	return
}
