// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// MagicTransitService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewMagicTransitService] method
// instead.
type MagicTransitService struct {
	Options         []option.RequestOption
	CfInterconnects *CfInterconnectService
	GRETunnels      *GRETunnelService
	IPSECTunnels    *IPSECTunnelService
	Routes          *RouteService
	Sites           *SiteService
}

// NewMagicTransitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagicTransitService(opts ...option.RequestOption) (r *MagicTransitService) {
	r = &MagicTransitService{}
	r.Options = opts
	r.CfInterconnects = NewCfInterconnectService(opts...)
	r.GRETunnels = NewGRETunnelService(opts...)
	r.IPSECTunnels = NewIPSECTunnelService(opts...)
	r.Routes = NewRouteService(opts...)
	r.Sites = NewSiteService(opts...)
	return
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39 string

const (
	UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39Reply   UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39 = "reply"
	UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39Request UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39 = "request"
)

func (r UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39Reply, UnnamedSchemaRef3b1a76a5e4a139b72ed7d93834773d39Request:
		return true
	}
	return false
}

// How frequent the health check is run. The default value is `mid`.
type UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5 string

const (
	UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5Low  UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5 = "low"
	UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5Mid  UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5 = "mid"
	UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5High UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5 = "high"
)

func (r UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5) IsKnown() bool {
	switch r {
	case UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5Low, UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5Mid, UnnamedSchemaRefEebdc868ce7f7ae92e23438caa84e7b5High:
		return true
	}
	return false
}
