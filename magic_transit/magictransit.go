// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// MagicTransitService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMagicTransitService] method instead.
type MagicTransitService struct {
	Options         []option.RequestOption
	CfInterconnects *CfInterconnectService
	GRETunnels      *GRETunnelService
	IPSECTunnels    *IPSECTunnelService
	Routes          *RouteService
	Sites           *SiteService
	Connectors      *ConnectorService
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
	r.Connectors = NewConnectorService(opts...)
	return
}

type HealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction HealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target string `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType `json:"type"`
	JSON healthCheckJSON `json:"-"`
}

// healthCheckJSON contains the JSON metadata for the struct [HealthCheck]
type healthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
// in health_check is ignored as the interface_address is used to send traffic into
// the tunnel.
type HealthCheckDirection string

const (
	HealthCheckDirectionUnidirectional HealthCheckDirection = "unidirectional"
	HealthCheckDirectionBidirectional  HealthCheckDirection = "bidirectional"
)

func (r HealthCheckDirection) IsKnown() bool {
	switch r {
	case HealthCheckDirectionUnidirectional, HealthCheckDirectionBidirectional:
		return true
	}
	return false
}

type HealthCheckParam struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel. Note in the case of bidirecitonal healthchecks, the target field
	// in health_check is ignored as the interface_address is used to send traffic into
	// the tunnel.
	Direction param.Field[HealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target.
	Target param.Field[string] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[HealthCheckType] `json:"type"`
}

func (r HealthCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
