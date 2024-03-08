// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustNetworkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustNetworkService] method
// instead.
type ZeroTrustNetworkService struct {
	Options         []option.RequestOption
	Routes          *ZeroTrustNetworkRouteService
	VirtualNetworks *ZeroTrustNetworkVirtualNetworkService
}

// NewZeroTrustNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustNetworkService(opts ...option.RequestOption) (r *ZeroTrustNetworkService) {
	r = &ZeroTrustNetworkService{}
	r.Options = opts
	r.Routes = NewZeroTrustNetworkRouteService(opts...)
	r.VirtualNetworks = NewZeroTrustNetworkVirtualNetworkService(opts...)
	return
}
