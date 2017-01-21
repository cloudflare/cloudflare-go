package virtualdns

import (
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// Cluster represents a Virtual DNS cluster.
type Cluster struct {
	DeprecateAnyRequests bool      `json:"deprecate_any_requests,omitempty"`
	ID                   string    `json:"id,omitempty"`
	MaximumCacheTTL      int       `json:"maximum_cache_ttl,omitempty"`
	MinimumCacheTTL      int       `json:"minimum_cache_ttl,omitempty"`
	ModifiedOn           time.Time `json:"modified_on,omitempty"`
	Name                 string    `json:"name,omitempty"`
	OriginIPs            []string  `json:"origin_ips,omitempty"`
	Ratelimit            int       `json:"ratelimit,omitempty"`
	VirtualDNSIPs        []string  `json:"virtual_dns_ips,omitempty"`
}

type VirtualDNSResponse struct {
	Result Cluster `json:"result"`
	cloudflare.Response
	cloudflare.ResultInfo `json:"result_info"`
}

type VirtualDNSListResponse struct {
	Result []Cluster `json:"result"`
	cloudflare.Response
	cloudflare.ResultInfo `json:"result_info"`
}
