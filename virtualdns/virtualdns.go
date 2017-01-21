// Package virtualdns supports managing Cloudflare's Virtual DNS service
// (https://www.cloudflare.com/dns/virtual-dns/).
package virtualdns

import (
	"encoding/json"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

type Client struct {
	client *cloudflare.API
	orgID  string
}

func NewClient(client *cloudflare.API) (*Client, error) {

	return nil, nil
}

func (*Client) SetOrg(organizationID string) error {

}

// GetClusters gets a list of the currently configured Virtual DNS clusters.
// API reference:
//   https://api.cloudflare.com/#...
//   GET /...
func (c *Client) GetClusters() ([]Cluster, error) {

	return []Cluster{}, nil
}

// GetClusterByID gets a Virtual DNS cluster by its ID.
// API reference:
//   https://api.cloudflare.com/#...
//   GET /...
func (c *Client) GetClusterByID(id string) (Cluster, error) {

	return Cluster{}, nil
}

// CreateCluster creates a Virtual DNS cluster against the provided backend IPs.
// API reference:
//   https://api.cloudflare.com/#...
//   POST /...
func (c *Client) CreateCluster(cluster Cluster) (Cluster, error) {
	// TODO(matt): provide an makePath(user) helper that constructs an
	// /organization/:org_id/... or /user/... path as appropriate.
	// 'user' and 'org' are enums that make it clearer than passing a boolean.
	// TODO(matt): Expose makeRequest -or- move it to cloudflare-go/internal/ - ?
	res, err := api.makeRequest("POST", makePath(c.orgID), cluster)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	response := &VirtualDNSResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// ModifyCluster modifies an existing Virtual DNS cluster.
// API reference:
//   https://api.cloudflare.com/#...
//   PATCH /...
func (c *Client) ModifyCluster(cluster Cluster) (Cluster, error) {

	return Cluster{}, nil
}

// DeleteCluster deletes an existing Virtual DNS cluster. Note that this will
// relinquish any Anycast IP addresses assigned to the cluster, and thus drop
// all queries to those addresses.
// API reference:
//   https://api.cloudflare.com/#...
//   DELETE /...
func (c *Client) DeleteCluster(cluster Cluster) (Cluster, error) {

	return Cluster{}, nil
}
