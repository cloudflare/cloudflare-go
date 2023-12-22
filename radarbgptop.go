// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarBgpTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarBgpTopService] method
// instead.
type RadarBgpTopService struct {
	Options  []option.RequestOption
	Ases     *RadarBgpTopAseService
	Prefixes *RadarBgpTopPrefixService
}

// NewRadarBgpTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpTopService(opts ...option.RequestOption) (r *RadarBgpTopService) {
	r = &RadarBgpTopService{}
	r.Options = opts
	r.Ases = NewRadarBgpTopAseService(opts...)
	r.Prefixes = NewRadarBgpTopPrefixService(opts...)
	return
}
