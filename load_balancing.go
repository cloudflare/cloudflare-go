package cloudflare

// Load Balancing API endpoints: https://www.cloudflare.com/load-balancing/

// Monitor represents a 'health check' configuration.
type Monitor struct {
	ID       string `json:"id,omitempty"`
	Created  string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
}

// MonitorResponse represents a response from the Monitors endpoint.
type MonitorResponse struct {
	Result []Monitor `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// MonitorListResponse repsents a list response from the Monitors endpoint.
type MonitorListResponse struct {
	Result []Monitor `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// Pool represents a 'pool' (group) of origin servers, each identified by their
// IP address or hostname. You can configure multiple Pools, and configure a
// failover priority (Pool A -> Pool B -> Pool C) as needed. If you're familiar
// with DNS terminology, think of a Pool as a DNS 'record set' that only
// returns healthy addresses.
type Pool struct{}

// PoolResponse represents a response from the Pools endpoint.
type PoolResponse struct{}

// PoolListResponse represents a list response from the Pools endpoint.
type PoolListResponse struct{}

// LB represents a Load Balancer: a DNS hostname and the load-balancing/failover/geo-routing policy for the associated Pools.
type LB struct{}

// LBResponse represents a list response from the Load Balancers endpoint.
type LBResponse struct{}

// LBListResponse represents a list response from the Load Balancers endpoint.
type LBListResponse struct{}
