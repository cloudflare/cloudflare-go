// Package virtualdns supports managing Cloudflare's Virtual DNS service
// (https://www.cloudflare.com/dns/virtual-dns/).
package virtualdns

import cloudflare "github.com/cloudflare/cloudflare-go"

type Client struct {
	client *cloudflare.API
	orgID  string
}

func NewClient(client *cloudflare.API) (*Client, error) {

	return nil, nil
}

// GetClusters gets a list of the currently configured Virtual DNS clusters.
// API reference:
//   https://api.cloudflare.com/#...
//   GET /...
func GetClusters() ([]Cluster, error) {

	return []Cluster{}, nil
}

// GetClusterByID gets a Virtual DNS cluster by its ID.
// API reference:
//   https://api.cloudflare.com/#...
//   GET /...
func GetClusterByID(id string) (Cluster, error) {

	return Cluster{}, nil
}

// CreateCluster creates a Virtual DNS cluster against the provided backend IPs.
// API reference:
//   https://api.cloudflare.com/#...
//   POST /...
func CreateCluster(cluster Cluster) (Cluster, error) {

	return Cluster{}, nil
}

// ModifyCluster modifies an existing Virtual DNS cluster.
// API reference:
//   https://api.cloudflare.com/#...
//   PATCH /...
func ModifyCluster(cluster Cluster) (Cluster, error) {

	return Cluster{}, nil
}

// DeleteCluster deletes an existing Virtual DNS cluster. Note that this will
// relinquish any Anycast IP addresses assigned to the cluster, and thus drop
// all queries to those addresses.
// API reference:
//   https://api.cloudflare.com/#...
//   DELETE /...
func DeleteCluster(cluster Cluster) (Cluster, error) {

	return Cluster{}, nil
}
