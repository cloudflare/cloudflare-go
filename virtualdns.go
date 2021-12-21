package cloudflare

import (
	"context"
	"fmt"
)

// NOTE: Everything in this file is deprecated. See `dns_firewall.go` instead.

// VirtualDNS represents a Virtual DNS configuration.
//
// Deprecated: Use DNSFirewallCluster instead.
type VirtualDNS struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	OriginIPs            []string `json:"origin_ips"`
	VirtualDNSIPs        []string `json:"virtual_dns_ips"`
	MinimumCacheTTL      uint     `json:"minimum_cache_ttl"`
	MaximumCacheTTL      uint     `json:"maximum_cache_ttl"`
	DeprecateAnyRequests bool     `json:"deprecate_any_requests"`
	ModifiedOn           string   `json:"modified_on"`
}

// VirtualDNSAnalyticsMetrics represents a group of aggregated Virtual DNS metrics.
//
// Deprecated: Use DNSFirewallAnalyticsMetrics instead.
type VirtualDNSAnalyticsMetrics DNSFirewallAnalyticsMetrics

// VirtualDNSAnalytics represents a set of aggregated Virtual DNS metrics.
//
// Deprecated: Use DNSFirewallAnalytics instead.
type VirtualDNSAnalytics struct {
	Totals VirtualDNSAnalyticsMetrics `json:"totals"`
	Min    VirtualDNSAnalyticsMetrics `json:"min"`
	Max    VirtualDNSAnalyticsMetrics `json:"max"`
}

// VirtualDNSUserAnalyticsOptions represents range and dimension selection on analytics endpoint
//
// Deprecated: Use DNSFirewallUserAnalyticsOptions instead.
type VirtualDNSUserAnalyticsOptions DNSFirewallUserAnalyticsOptions

// VirtualDNSResponse represents a Virtual DNS response.
//
// Deprecated: This internal type will be removed in the future.
type VirtualDNSResponse struct {
	Response
	Result *VirtualDNS `json:"result"`
}

// VirtualDNSListResponse represents an array of Virtual DNS responses.
//
// Deprecated: This internal type will be removed in the future.
type VirtualDNSListResponse struct {
	Response
	Result []*VirtualDNS `json:"result"`
}

// VirtualDNSAnalyticsResponse represents a Virtual DNS analytics response.
//
// Deprecated: This internal type will be removed in the future.
type VirtualDNSAnalyticsResponse struct {
	Response
	Result VirtualDNSAnalytics `json:"result"`
}

// CreateVirtualDNS creates a new Virtual DNS cluster.
//
// Deprecated: Use CreateDNSFirewallCluster instead.
func (api *API) CreateVirtualDNS(ctx context.Context, v *VirtualDNS) (*VirtualDNS, error) {
	if v == nil {
		return nil, fmt.Errorf("cluster must not be nil")
	}

	res, err := api.CreateDNSFirewallCluster(ctx, v.vdnsUpgrade())
	return res.vdnsDowngrade(), err
}

// VirtualDNS fetches a single virtual DNS cluster.
//
// Deprecated: Use DNSFirewallCluster instead.
func (api *API) VirtualDNS(ctx context.Context, virtualDNSID string) (*VirtualDNS, error) {
	res, err := api.DNSFirewallCluster(ctx, virtualDNSID)
	return res.vdnsDowngrade(), err
}

// ListVirtualDNS lists the virtual DNS clusters associated with an account.
//
// Deprecated: Use ListDNSFirewallClusters instead.
func (api *API) ListVirtualDNS(ctx context.Context) ([]*VirtualDNS, error) {
	res, err := api.ListDNSFirewallClusters(ctx)
	if res == nil {
		return nil, err
	}

	clusters := make([]*VirtualDNS, 0, len(res))
	for _, cluster := range res {
		clusters = append(clusters, cluster.vdnsDowngrade())
	}

	return clusters, err
}

// UpdateVirtualDNS updates a Virtual DNS cluster.
//
// Deprecated: Use UpdateDNSFirewallCluster instead.
func (api *API) UpdateVirtualDNS(ctx context.Context, virtualDNSID string, vv VirtualDNS) error {
	return api.UpdateDNSFirewallCluster(ctx, virtualDNSID, vv.vdnsUpgrade())
}

// DeleteVirtualDNS deletes a Virtual DNS cluster. Note that this cannot be
// undone, and will stop all traffic to that cluster.
//
// Deprecated: Use DeleteDNSFirewallCluster instead.
func (api *API) DeleteVirtualDNS(ctx context.Context, virtualDNSID string) error {
	return api.DeleteDNSFirewallCluster(ctx, virtualDNSID)
}

// VirtualDNSUserAnalytics retrieves analytics report for a specified dimension and time range
//
// Deprecated: Use DNSFirewallUserAnalytics instead.
func (api *API) VirtualDNSUserAnalytics(ctx context.Context, virtualDNSID string, o VirtualDNSUserAnalyticsOptions) (VirtualDNSAnalytics, error) {
	res, err := api.DNSFirewallUserAnalytics(ctx, virtualDNSID, DNSFirewallUserAnalyticsOptions(o))
	return res.vdnsDowngrade(), err
}

// --- Compatibility helper functions ---

func (v VirtualDNS) vdnsUpgrade() DNSFirewallCluster {
	return DNSFirewallCluster{
		ID:                   v.ID,
		Name:                 v.Name,
		OriginIPs:            v.OriginIPs,
		DNSFirewallIPs:       v.VirtualDNSIPs,
		MinimumCacheTTL:      v.MinimumCacheTTL,
		MaximumCacheTTL:      v.MaximumCacheTTL,
		DeprecateAnyRequests: v.DeprecateAnyRequests,
		ModifiedOn:           v.ModifiedOn,
	}
}

func (v *DNSFirewallCluster) vdnsDowngrade() *VirtualDNS {
	if v == nil {
		return nil
	}

	return &VirtualDNS{
		ID:                   v.ID,
		Name:                 v.Name,
		OriginIPs:            v.OriginIPs,
		VirtualDNSIPs:        v.DNSFirewallIPs,
		MinimumCacheTTL:      v.MinimumCacheTTL,
		MaximumCacheTTL:      v.MaximumCacheTTL,
		DeprecateAnyRequests: v.DeprecateAnyRequests,
		ModifiedOn:           v.ModifiedOn,
	}
}

func (v DNSFirewallAnalytics) vdnsDowngrade() VirtualDNSAnalytics {
	return VirtualDNSAnalytics{
		Totals: VirtualDNSAnalyticsMetrics(v.Totals),
		Min:    VirtualDNSAnalyticsMetrics(v.Min),
		Max:    VirtualDNSAnalyticsMetrics(v.Max),
	}
}
