// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"github.com/cloudflare/cloudflare-go/v3/option"
)

// MagicTransitService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMagicTransitService] method instead.
type MagicTransitService struct {
	Options         []option.RequestOption
	Apps            *AppService
	CfInterconnects *CfInterconnectService
	GRETunnels      *GRETunnelService
	IPSECTunnels    *IPSECTunnelService
	Routes          *RouteService
	Sites           *SiteService
	Connectors      *ConnectorService
	PCAPs           *PCAPService
}

// NewMagicTransitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMagicTransitService(opts ...option.RequestOption) (r *MagicTransitService) {
	r = &MagicTransitService{}
	r.Options = opts
	r.Apps = NewAppService(opts...)
	r.CfInterconnects = NewCfInterconnectService(opts...)
	r.GRETunnels = NewGRETunnelService(opts...)
	r.IPSECTunnels = NewIPSECTunnelService(opts...)
	r.Routes = NewRouteService(opts...)
	r.Sites = NewSiteService(opts...)
	r.Connectors = NewConnectorService(opts...)
	r.PCAPs = NewPCAPService(opts...)
	return
}

// How frequent the health check is run. The default value is `mid`.
type HealthCheckRate string

const (
	HealthCheckRateLow  HealthCheckRate = "low"
	HealthCheckRateMid  HealthCheckRate = "mid"
	HealthCheckRateHigh HealthCheckRate = "high"
)

func (r HealthCheckRate) IsKnown() bool {
	switch r {
	case HealthCheckRateLow, HealthCheckRateMid, HealthCheckRateHigh:
		return true
	}
	return false
}

// The type of healthcheck to run, reply or request. The default value is `reply`.
type HealthCheckType string

const (
	HealthCheckTypeReply   HealthCheckType = "reply"
	HealthCheckTypeRequest HealthCheckType = "request"
)

func (r HealthCheckType) IsKnown() bool {
	switch r {
	case HealthCheckTypeReply, HealthCheckTypeRequest:
		return true
	}
	return false
}
