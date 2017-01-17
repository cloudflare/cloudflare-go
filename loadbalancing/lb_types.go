package loadbalancer

import (
	"net/http"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// Monitor represents a 'health check' configuration. A Monitor will only
// actively health check origin servers once applied to a Pool.
type Monitor struct {
	CreatedOn       time.Time   `json:"created_on"`
	Description     string      `json:"description"`
	ExpectedBody    string      `json:"expected_body"`
	ExpectedCodes   string      `json:"expected_codes"`
	FollowRedirects bool        `json:"follow_redirects"`
	Header          http.Header `json:"header"`
	ID              string      `json:"id"`
	Interval        int         `json:"interval"`
	Method          string      `json:"method"`
	ModifiedOn      time.Time   `json:"modified_on"`
	Path            string      `json:"path"`
	Retries         int         `json:"retries"`
	Timeout         int         `json:"timeout"`
	Type            string      `json:"type"`
}

// MonitorResponse represents a response from the Monitors endpoint.
type MonitorResponse struct {
	Result []Monitor `json:"result"`
	cloudflare.Response
	cloudflare.ResultInfo `json:"result_info"`
}

// MonitorListResponse repsents a list response from the Monitors endpoint.
type MonitorListResponse struct {
	Result []Monitor `json:"result"`
	cloudflare.Response
	cloudflare.ResultInfo `json:"result_info"`
}

// Pool represents a 'pool' (group) of origin servers, each identified by their
// IP address or hostname. You can configure multiple Pools, and configure a
// failover priority (Pool A -> Pool B -> Pool C) as needed. If you're familiar
// with DNS terminology, think of a Pool as a DNS 'record set' that only
// returns healthy addresses.
type Pool struct {
	CreatedOn         time.Time `json:"created_on"`
	Description       string    `json:"description"`
	Enabled           bool      `json:"enabled"`
	ID                string    `json:"id"`
	MinimumOrigins    int       `json:"minimum_origins"`
	ModifiedOn        time.Time `json:"modified_on"`
	Monitor           string    `json:"monitor"`
	Name              string    `json:"name"`
	NotificationEmail string    `json:"notification_email"`
	Origins           []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Enabled bool   `json:"enabled"`
	} `json:"origins"`
}

// PoolResponse represents a response from the Pools endpoint.
type PoolResponse struct{}

// PoolListResponse represents a list response from the Pools endpoint.
type PoolListResponse struct{}

// LB represents a Load Balancer: a DNS hostname and the load-balancing/failover/geo-routing policy for the associated Pools.
type LB struct {
	CreatedOn    time.Time `json:"created_on"`
	DefaultPools []string  `json:"default_pools"`
	Description  string    `json:"description"`
	FallbackPool string    `json:"fallback_pool"`
	ID           string    `json:"id"`
	ModifiedOn   time.Time `json:"modified_on"`
	Name         string    `json:"name"`
	// PopPools are an Enterprise-only feature for mapping specific Cloudflare
	// data centers to pools of origin servers.
	POPPools    map[string][]string `json:"pop_pools"`
	Proxied     bool                `json:"proxied"`
	RegionPools map[string][]string `json:"region_pools"`
	TTL         int                 `json:"ttl"`
}

// LBResponse represents a list response from the Load Balancers endpoint.
type LBResponse struct{}

// LBListResponse represents a list response from the Load Balancers endpoint.
type LBListResponse struct{}
