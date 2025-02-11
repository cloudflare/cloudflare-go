// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/option"
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

type HealthCheckParam struct {
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[HealthCheckTargetUnionParam] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[HealthCheckType] `json:"type"`
}

func (r HealthCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by [magic_transit.HealthCheckTargetMagicHealthCheckTargetParam],
// [shared.UnionString].
type HealthCheckTargetUnionParam interface {
	ImplementsHealthCheckTargetUnionParam()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type HealthCheckTargetMagicHealthCheckTargetParam struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r HealthCheckTargetMagicHealthCheckTargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r HealthCheckTargetMagicHealthCheckTargetParam) ImplementsHealthCheckTargetUnionParam() {}

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
