// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMagicService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMagicService] method
// instead.
type AccountMagicService struct {
	Options         []option.RequestOption
	CfInterconnects *AccountMagicCfInterconnectService
	GreTunnels      *AccountMagicGreTunnelService
	IpsecTunnels    *AccountMagicIpsecTunnelService
	Routes          *AccountMagicRouteService
}

// NewAccountMagicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMagicService(opts ...option.RequestOption) (r *AccountMagicService) {
	r = &AccountMagicService{}
	r.Options = opts
	r.CfInterconnects = NewAccountMagicCfInterconnectService(opts...)
	r.GreTunnels = NewAccountMagicGreTunnelService(opts...)
	r.IpsecTunnels = NewAccountMagicIpsecTunnelService(opts...)
	r.Routes = NewAccountMagicRouteService(opts...)
	return
}
