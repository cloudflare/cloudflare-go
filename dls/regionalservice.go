// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls

import (
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// RegionalServiceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalServiceService] method instead.
type RegionalServiceService struct {
	Options        []option.RequestOption
	PrefixBindings *RegionalServicePrefixBindingService
}

// NewRegionalServiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegionalServiceService(opts ...option.RequestOption) (r *RegionalServiceService) {
	r = &RegionalServiceService{}
	r.Options = opts
	r.PrefixBindings = NewRegionalServicePrefixBindingService(opts...)
	return
}
