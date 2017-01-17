// Package loadbalancer provides an interface for configuring Cloudflare's Load Balancing
// API endpoints: https://www.cloudflare.com/load-balancing/
package loadbalanacer

import cloudflare "github.com/cloudflare/cloudflare-go"

// Client ...
type Client struct{ apiClient *cloudflare.API }

// NewClient ...
func NewClient(api *cloudflare.API, organizationID string) (*Client, error) {
	return nil, nil
}

// GetLoadBalancers returns a list of the currently configured Load Balancers for the given zone ID.
// API reference:
// 	https://api.cloudflare.com/#load-balancing-#get-load-balancer
//	GET /zone/:zone_id/load_balancers
func (c *Client) GetLoadBalancers(zoneID string) ([]LB, error) {

	return LBListResponse{}, nil
}

// GetLoadBalancerByID returns a single Load Balancer by its ID.
func (c *Client) GetLoadBalancerByID(id string) (LB, error) {

	return LBResponse{}, nil
}

// CreateLoadBalancer creates a Load Balancer.
func (c *Client) CreateLoadBalancer(lb LB) (LB, error) {

	return LB, nil
}

// UpdateLoadBalancer updates an existing Load Balancer object. The existing Load Balancer is
// replaced, not patched.
func (c *Client) UpdateLoadBalancer(lb LB) error {

	return LBResponse{}, nil
}

// GetPools returns a list of the currently configured Pools.
func (c *Client) GetPools() ([]Pool, error) {

	return PoolListResponse{}, nil
}

// GetPoolByID returns a single Pool by its ID.
func (c *Client) GetPoolByID(id string) (Pool, error) {

	return PoolResponse, nil
}

// CreatePool creates a Pool.
func (c *Client) CreatePool(pool Pool) error {

	return PoolResponse, nil
}

// UpdatePool updates an existing Pool object. The existing Pool is replaced,
// not patched.
func (c *Client) UpdatePool(pool Pool) error {

	return PoolResponse, nil
}

// GetMonitors returns a list of the currently configured Monitors.
func (c *Client) GetMonitors() (MonitorListResponse, error) {

	return MonitorListResponse{}, nil
}

// GetMonitorByID returns a single Monitor by its ID.
func (c *Client) GetMonitorByID(id string) (MonitorResponse, error) {

	return MonitorResponse, nil
}

// CreateMonitor creates a Monitor.
func (c *Client) CreateMonitor(monitor Monitor) (MonitorResponse, error) {

	return MonitorResponse, nil
}

// UpdateMonitor updates an existing Monitor object. The existing Monitor is
// replaced, not patched.
func (c *Client) UpdateMonitor(monitor Monitor) (MonitorResponse, error) {

	return MonitorResponse, nil
}
