// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZeroTrustService] method instead.
type ZeroTrustService struct {
	Options              []option.RequestOption
	IdentityProviders    *ZeroTrustIdentityProviderService
	Organizations        *ZeroTrustOrganizationService
	Seats                *ZeroTrustSeatService
	Access               *ZeroTrustAccessService
	DEX                  *ZeroTrustDEXService
	Tunnels              *ZeroTrustTunnelService
	ConnectivitySettings *ZeroTrustConnectivitySettingService
	DLP                  *ZeroTrustDLPService
	Gateway              *ZeroTrustGatewayService
	Networks             *ZeroTrustNetworkService
}

// NewZeroTrustService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustService(opts ...option.RequestOption) (r *ZeroTrustService) {
	r = &ZeroTrustService{}
	r.Options = opts
	r.IdentityProviders = NewZeroTrustIdentityProviderService(opts...)
	r.Organizations = NewZeroTrustOrganizationService(opts...)
	r.Seats = NewZeroTrustSeatService(opts...)
	r.Access = NewZeroTrustAccessService(opts...)
	r.DEX = NewZeroTrustDEXService(opts...)
	r.Tunnels = NewZeroTrustTunnelService(opts...)
	r.ConnectivitySettings = NewZeroTrustConnectivitySettingService(opts...)
	r.DLP = NewZeroTrustDLPService(opts...)
	r.Gateway = NewZeroTrustGatewayService(opts...)
	r.Networks = NewZeroTrustNetworkService(opts...)
	return
}
